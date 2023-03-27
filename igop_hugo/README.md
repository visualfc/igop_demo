```
git clone https://github.com/gohugoio/hugo
cd hugo
go mod tidy
cd ..
igop_hugo ./hugo new site quickstart
cd quickstart
git init
git submodule add https://github.com/theNewDynamic/gohugo-theme-ananke themes/ananke
echo "theme = 'ananke'" >> config.toml
igop_hugo ../hugo server
```