# examples

## Hertz examples

```bash
go install github.com/cloudwego/hertz/cmd/hz@latest
```

## Kitex examples

```bash
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```

## MacOS

```bash
echo "**/.DS_Store" >> ~/.gitignore_global
git config --global core.excludesfile ~/.gitignore_global
```

## How to fork

```bash
git remote remove origin
git remote add origin <forkUrl>
git checkout -b fix
```

- biz: business
- dal: data access layer
- mw: middleware
