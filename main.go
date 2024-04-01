package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {

	cfg := struct {
		nodeURL string
	}{}

	app := &cli.App{
		Name:                 "rpcrequest",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "node-url",
				Required:    true,
				Usage:       "node url",
				EnvVars:     []string{"NODE_URL"},
				Destination: &cfg.nodeURL,
			},
		},
		Action: func(c *cli.Context) error {

			if !c.Args().Present() {
				return fmt.Errorf("method is required")
			}

			method := c.Args().First()

			rawParams := c.Args().Slice()[1:]

			params := []any{}
			for _, p := range rawParams {
				v, err := transform(p)
				if err != nil {
					return fmt.Errorf("failed to transform param %q: %w", p, err)
				}
				params = append(params, v)
			}

			rpcRequest := rpcRequest{
				ID:      0,
				Jsonrpc: "2.0",
				Method:  method,
				Params:  params,
			}

			d, err := json.Marshal(rpcRequest)
			if err != nil {
				return fmt.Errorf("failed to marshal rpc request: %w", err)
			}

			req, err := http.NewRequest("POST", cfg.nodeURL, bytes.NewReader(d))
			if err != nil {
				return fmt.Errorf("failed to create request: %w", err)
			}

			req.Header.Set("Content-Type", "application/json")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return fmt.Errorf("failed to send request: %w", err)
			}

			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				d, _ := io.ReadAll(res.Body)
				return fmt.Errorf("unexpected status %s: %s", res.Status, string(d))
			}

			resp := rpcResponse{}

			err = json.NewDecoder(res.Body).Decode(&resp)
			if err != nil {
				return fmt.Errorf("failed to decode response: %w", err)
			}

			if len(resp.Error) > 0 {
				return fmt.Errorf("rpc error: %s", string(resp.Error))
			}

			fmt.Println(string(resp.Result))

			return nil

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func transform(v string) (any, error) {
	switch {
	case strings.HasPrefix(v, "bool:"):
		return v[5:] == "true", nil
	case strings.HasPrefix(v, "asHex:"):
		iv, err := strconv.ParseInt(v[6:], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int: %w", err)
		}
		return fmt.Sprintf("0x%x", iv), nil
	case strings.HasPrefix(v, "int:"):
		iv, err := strconv.ParseInt(v[4:], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse int: %w", err)
		}
		return iv, nil
	case strings.HasPrefix(v, "json:"):
		var res any
		err := json.Unmarshal([]byte(v[5:]), &res)
		if err != nil {
			return nil, fmt.Errorf("failed to parse json: %w", err)
		}
		return res, nil
	default:
		return v, nil
	}

}

type rpcRequest struct {
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      any    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type rpcResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   json.RawMessage `json:"error"`
	ID      any             `json:"id"`
}
