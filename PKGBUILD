# Maintainer: Lukasz Magiera <me@magik6k.net>

pkgname=tappipe
pkgver=1.0
pkgrel=1
pkgdesc="Simple utility allowing to create TAP interface using stdin/out as data channel"
arch=('x86_64' 'i686')
url="https://github.com/magik6k/tapPipe"
license=('MIT')
makedepends=('go')
options=('!strip' '!emptydirs')

build() {
  cd ..
  go build
}

package() {
  cd ..
  install -Dm755 "tapPipe" "$pkgdir/usr/bin/$pkgname"
}

# vim:set ts=2 sw=2 et:
