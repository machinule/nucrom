# nucrom

## Installation

1. Install msys2: https://msys2.github.io/
2. In the msys2 shell:
  `pacman -S patch`
  `pacman -S mingw-w64-x86_64-go`
  `pacman -S git`
  `git clone git://github.com/Alexpux/MINGW-packages.git`
  `cd MINGW-packages/mingw-w64-go`
  `makepkg-mingw`
  `pacman -U mingw-w64-x86_64-go-1.6-1-any.pkg.tar.xz`
  `pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-SDL2{,_mixer,_image,_ttf}`
3. In the mingw64 shell:
  `mkdir gowork`
  `echo 'export GOPATH=$HOME/gowork' >> ~/.bashrc'`
  `echo 'export GOROOT=/mingw64/lib/go' >> ~/.bashrc'`
  `source ~/.bashrc`
  `cd gowork`
  `go get -v github.com/veandco/go-sdl2/sdl{,_mixer,_image,_ttf}`
  `go get -v github.com/go-gl/gl/v{3.2,3.3,4.1,4.4,4.5}-{core,compatibility}/gl`
  `go get -v github.com/go-gl/gl/v3.3-core/gl`
  `cd src/github.com/`
  `mkdir machinule`
  `cd machinule`
  `git clone git://github.com/machinule/nucrom.git`
  