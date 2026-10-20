package bench

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/parquet-go/parquet-go"
)

// Dataset is an open dataset the benchmark can download.
type Dataset struct {
	Name, Title, License, URL string
}

// Datasets lists the sources for Scenario.Data.
var Datasets = []Dataset{
	{"oasst2", "OpenAssistant Conversations 2 — human prompts and assistant replies, many languages", "Apache-2.0",
		"https://huggingface.co/datasets/OpenAssistant/oasst2/resolve/main/2023-11-05_oasst2_ready.messages.jsonl.gz"},
	{"swe-gym", "SWE-Gym OpenHands trajectories — a coding agent's actions and tool outputs", "MIT",
		"https://huggingface.co/datasets/SWE-Gym/OpenHands-SFT-Trajectories/resolve/main/data/train.success.oss-00000-of-00001.parquet"},
}

// CacheDir is where datasets are kept (P2PACT_BENCH_CACHE overrides).
func CacheDir() string {
	if d := os.Getenv("P2PACT_BENCH_CACHE"); d != "" {
		return d
	}
	d, err := os.UserCacheDir()
	if err != nil {
		d = os.TempDir()
	}
	return filepath.Join(d, "p2pact-bench")
}

// LoadData returns the texts for a scenario's Data field and a line crediting the source.
func LoadData(spec string, log io.Writer) ([]string, string, error) {
	parts := strings.Split(spec, ":")
	switch parts[0] {
	case "synthetic":
		return nil, "generated text (Zipf-distributed words)", nil
	case "file":
		path := strings.TrimPrefix(spec, "file:")
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, "", err
		}
		var msgs []string
		if err := json.Unmarshal(b, &msgs); err != nil {
			return nil, "", fmt.Errorf("%s: want a JSON array of strings: %w", path, err)
		}
		return msgs, "your own messages (" + filepath.Base(path) + ")", nil
	case "oasst2":
		if len(parts) < 2 || (parts[1] != "human" && parts[1] != "assistant") {
			return nil, "", fmt.Errorf("want oasst2:human or oasst2:assistant[:lang]")
		}
		lang := ""
		if len(parts) > 2 {
			lang = parts[2]
		}
		f, err := fetch(Datasets[0], log)
		if err != nil {
			return nil, "", err
		}
		msgs, err := readOASST(f, map[string]string{"human": "prompter", "assistant": "assistant"}[parts[1]], lang)
		credit := "OpenAssistant oasst2 (Apache-2.0), " + parts[1]
		if lang != "" {
			credit += ", language " + lang
		}
		return msgs, credit, err
	case "swe-gym":
		if len(parts) < 2 || (parts[1] != "agent" && parts[1] != "tool") {
			return nil, "", fmt.Errorf("want swe-gym:agent or swe-gym:tool")
		}
		f, err := fetch(Datasets[1], log)
		if err != nil {
			return nil, "", err
		}
		msgs, err := readSWEGym(f, parts[1])
		return msgs, "SWE-Gym OpenHands trajectories (MIT), " + parts[1] + " messages", err
	}
	return nil, "", fmt.Errorf("unknown data %q: synthetic, oasst2:human|assistant[:lang], swe-gym:agent|tool, file:<path>", spec)
}

// fetch downloads a dataset file once into the cache.
func fetch(d Dataset, log io.Writer) (string, error) {
	dir := CacheDir()
	path := filepath.Join(dir, filepath.Base(d.URL))
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	fmt.Fprintf(log, "downloading %s (%s) to %s\n", d.Name, d.License, path)
	resp, err := http.Get(d.URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%s: %s", d.URL, resp.Status)
	}
	tmp := path + ".part"
	f, err := os.Create(tmp)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	return path, os.Rename(tmp, path)
}

func readOASST(path, role, lang string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(gz)
	sc.Buffer(make([]byte, 1<<20), 64<<20)
	var out []string
	for sc.Scan() {
		var m struct {
			Role, Text, Lang string
		}
		if json.Unmarshal(sc.Bytes(), &m) != nil || m.Role != role || (lang != "" && m.Lang != lang) {
			continue
		}
		if m.Text != "" {
			out = append(out, m.Text)
		}
	}
	return out, sc.Err()
}

func readSWEGym(path, kind string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	type msg struct {
		Role    string `parquet:"role"`
		Content string `parquet:"content"`
	}
	type row struct {
		Messages []msg `parquet:"messages,list"`
	}
	r := parquet.NewGenericReader[row](f)
	defer r.Close()
	var out []string
	buf := make([]row, 64)
	for {
		n, err := r.Read(buf)
		for _, rw := range buf[:n] {
			firstUser := true
			for _, m := range rw.Messages {
				switch {
				case kind == "agent" && m.Role == "assistant":
					out = append(out, m.Content)
				case kind == "tool" && m.Role == "user":
					if firstUser { // the task itself, not a tool output
						firstUser = false
						continue
					}
					out = append(out, m.Content)
				}
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}
