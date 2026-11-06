package main
import("os";"path/filepath";"testing")
func TestSnapshotAndDiff(t *testing.T){dir:=t.TempDir();path:=filepath.Join(dir,"config");os.WriteFile(path,[]byte("a"),0600);a,e:=snapshot(dir);if e!=nil{t.Fatal(e)};os.WriteFile(path,[]byte("b"),0600);b,_:=snapshot(dir);if len(diff(a,b))!=1||diff(a,b)[0]!="CHANGED config"{t.Fatal(diff(a,b))};os.Remove(path);c,_:=snapshot(dir);if diff(b,c)[0]!="DELETED config"{t.Fatal(diff(b,c))}}
