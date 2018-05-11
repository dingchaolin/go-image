# go-image

## mac环境搭建

### libvips 安装
- 下载 https://github.com/jcupitt/libvips/releases/download/v8.6.3/vips-8.6.3.tar.gz
- 解压
- ./configure
- make
- sudo make install

## centos 安装
- 下载 wget https://github.com/jcupitt/libvips/releases/download/v8.6.3/vips-8.6.3.tar.gz
- tar -zxf vips-8.6.3.tar.gz
- mv vips-8.6.3.tar vips
- ./configure
- make install

#### centos 安装过程出现的问题 可能需要安装一下包
- yum install libpng-devel
- yum install glib2-devel
- yum -y install expat-devel

## centos 安装go
- yum install go

## Package vips was not found in the pkg-config search path
-  export PKG_CONFIG_PATH=/usr/lib64/pkgconfig:/usr/share/pkgconfig:/usr/lib/pkgconfig:/usr/local/lib/pkgconfig:/usr/local/share/pkgconfig

## mac 和 centos 在各自的环境中编译