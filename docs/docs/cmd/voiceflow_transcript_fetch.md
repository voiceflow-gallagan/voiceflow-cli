# voiceflow transcript fetch

Fetch one transcripts from a project

```
voiceflow transcript fetch [flags]
```

## Options

```
  -a, --agent-id string           Voiceflow Agent ID (required)
  -h, --help                      help for fetch
  -d, --output-directory string   Output directory to save the transcripts. Default is ./output (optional) (default "./output")
  -t, --transcript-id string      Voiceflow Transcript ID (required)
```

## Options inherited from parent commands

```
  -o, --output-format string         Output Format. Options: text, json. Default: text (optional) (default "text")
  -u, --skip-update-check            Skip the check for updates check run before every command (optional)
  -v, --verbose                      verbose error output (with stack trace) (optional)
  -x, --voiceflow-api-key string     Voiceflow API Key (optional)
  -b, --voiceflow-subdomain string   Voiceflow Base URL (optional). Default: empty
```

## See also

* [voiceflow transcript](/cmd/voiceflow_transcript/)	 - Actions on transcripts
