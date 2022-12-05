<h1 align="center">Jwt Auth Golang</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/maurodesouza/jwt-auth-golang?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/maurodesouza/jwt-auth-golang?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/maurodesouza/jwt-auth-golang?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/maurodesouza/jwt-auth-golang?color=56BEB8">
</p>

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#rocket-main-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/maurodesouza" target="_blank">Author</a>
</p>

<br>

## :dart: About ##

I create this project to improve my go lang skills, I had a little help from Robby with amazing this [tutorial](https://www.youtube.com/watch?v=ma7rUS_vW9M), thanks Robby ❤

## :rocket: Main Technologies ##

<a href="https://go.dev">
  <img width="50" title="Go Lang" alt="Go Logo" src="https://github.com/maurodesouza/maurodesouza/raw/master/assets/golang-logo.svg">
</a> &#xa0; &#xa0;

<a href="https://www.docker.com">
  <img width="50" title="Docker" alt="Docker Logo" src="./.github/assets/docker.svg">
</a> &#xa0; &#xa0;

###

<details>
  <summary>See more</summary>

  ###

- [Gin](https://gin-gonic.com)
- [GORM](https://gorm.io)
- [Go JWT](https://github.com/golang-jwt/jwt)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Dot Env](https://github.com/joho/godotenv)
- [Compile Daemon](https://github.com/githubnemo/CompileDaemon)

</details>

## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Docker](https://www.docker.com) installed.

## :checkered_flag: Starting ##

```bash
# Clone this project
$ git clone https://github.com/maurodesouza/jwt-auth-golang

# Access
$ cd jwt-auth-golang

# Start containers
$ docker-compose up -d

# Enter the server container
$ docker-compose exec -it server sh

# Start dev server
$ CompileDaemon -command="./jwt-auth-golang"

# The server will initialize in the <http://localhost:3333>
```

There's a insomia-endpoint.json file with all usable endpoints

## :memo: License ##

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.


Made with :heart: by <a href="https://github.com/maurodesouza" target="_blank">Mauro de Souza</a>

&#xa0;

<a href="#top">Back to top</a>
