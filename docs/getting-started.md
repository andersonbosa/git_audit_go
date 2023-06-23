# docs/getting-started.md

## Instalation

```bash
# download the binary
wget -L https://github.com/andersonbosa/git_audit_go/raw/main/git_audit_go/git_audit_go -o $HOME/.local/bin/git_audit_go

# give execution permission
chmod +x $HOME/.local/bin/git_audit_go
```

## Usage

```bash
# enter in a versioned repository
cd /tmp/railsgoat

# execute
git_audit_go

# inspect the output
open ./output.csv
```