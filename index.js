const go = new Go();

WebAssembly.instantiateStreaming(fetch("vita.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
