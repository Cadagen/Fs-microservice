outdir="./build"
name="cadagen-fs"
entrypoint="./src/main.go"

echo "Building $entrypoint into $outdir/$name..."

rm -r $outdir &> /dev/null || true

go build -o $outdir/$name $entrypoint

echo "Done."