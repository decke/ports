--- vendor/github.com/docker/libnetwork/drivers/bridge/setup_ipv4.go.orig	2020-09-04 14:54:57 UTC
+++ vendor/github.com/docker/libnetwork/drivers/bridge/setup_ipv4.go
@@ -1,3 +1,5 @@
+// +build linux
+
 package bridge
 
 import (
