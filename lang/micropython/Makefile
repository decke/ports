PORTNAME=	micropython
PORTVERSION=	1.14
CATEGORIES=	lang python
MASTER_SITES=	http://micropython.org/resources/source/

MAINTAINER=	ports@FreeBSD.org
COMMENT=	Implementation of the Python language for microcontrollers

LICENSE=	MIT
LICENSE_FILE=	${WRKDIR}/${DISTNAME}/LICENSE

BUILD_DEPENDS=	bash:shells/bash \
		gcc10:lang/gcc10
RUN_DEPENDS=	bash:shells/bash \
		gcc10:lang/gcc10
LIB_DEPENDS=	libffi.so:devel/libffi

USES=		gmake pkgconfig python:build tar:xz

PLIST_FILES=	bin/micropython

WRKSRC_SUBDIR=	ports/unix
MAKE_ARGS=	V=1 PYTHON=${PYTHON_CMD} CC=gcc10
ALL_TARGET=	axtls micropython
TEST_TARGET=	test

.include <bsd.port.mk>
