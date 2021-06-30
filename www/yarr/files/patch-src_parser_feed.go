--- src/parser/feed.go.orig	2021-06-07 09:07:20 UTC
+++ src/parser/feed.go
@@ -18,6 +18,11 @@ type processor func(r io.Reader) (*Feed, error)
 func sniff(lookup string) (string, processor) {
 	lookup = strings.TrimSpace(lookup)
 	lookup = strings.TrimLeft(lookup, "\x00\xEF\xBB\xBF\xFE\xFF")
+
+	if len(lookup) < 1 {
+		return "", nil
+	}
+
 	switch lookup[0] {
 	case '<':
 		decoder := xmlDecoder(strings.NewReader(lookup))
