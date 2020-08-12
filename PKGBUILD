# Maintainer: Ron B.S <ronthecookie0101 on gmail --OR-- me AT ronthecookie DOT me >

pkgname=golaunch
# https://gitlab.com/diamondburned/6cord/releases
pkgver=0.2
pkgrel=1
pkgdesc='a simple .desktop launcher written in go.'
arch=('x86_64')
url="https://github.com/hen6003/golaunch"
license=('MIT')
makedepends=('go' 'git')
source=("https://github.com/hen6003/golaunch/archive/master.tar.gz")
md5sums=("c343bd8adbaa6b1b292330ca0518cbf5")

build() {
  cd "$pkgname"-master
  go build \
    -gcflags "all=-trimpath=$PWD" \
    -asmflags "all=-trimpath=$PWD" \
    -ldflags "-extldflags $LDFLAGS" \
    -o $pkgname .
}

package() {
  cd "$pkgname"-master
  install -Dm755 $pkgname "$pkgdir"/usr/bin/"$pkgname"
}
