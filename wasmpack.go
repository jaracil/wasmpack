package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
)

var base64Code = `(function(r){if(typeof exports==="object"&&typeof module!=="undefined"){module.exports=r()}else if(typeof define==="function"&&define.amd){define([],r)}else{var e;if(typeof window!=="undefined"){e=window}else if(typeof global!=="undefined"){e=global}else if(typeof self!=="undefined"){e=self}else{e=this}e.base64js=r()}})(function(){var r,e,n;return function(){function r(e,n,t){function o(f,i){if(!n[f]){if(!e[f]){var u="function"==typeof require&&require;if(!i&&u)return u(f,!0);if(a)return a(f,!0);var v=new Error("Cannot find module '"+f+"'");throw v.code="MODULE_NOT_FOUND",v}var d=n[f]={exports:{}};e[f][0].call(d.exports,function(r){var n=e[f][1][r];return o(n||r)},d,d.exports,r,e,n,t)}return n[f].exports}for(var a="function"==typeof require&&require,f=0;f<t.length;f++)o(t[f]);return o}return r}()({"/":[function(r,e,n){"use strict";n.byteLength=d;n.toByteArray=h;n.fromByteArray=p;var t=[];var o=[];var a=typeof Uint8Array!=="undefined"?Uint8Array:Array;var f="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";for(var i=0,u=f.length;i<u;++i){t[i]=f[i];o[f.charCodeAt(i)]=i}o["-".charCodeAt(0)]=62;o["_".charCodeAt(0)]=63;function v(r){var e=r.length;if(e%4>0){throw new Error("Invalid string. Length must be a multiple of 4")}var n=r.indexOf("=");if(n===-1)n=e;var t=n===e?0:4-n%4;return[n,t]}function d(r){var e=v(r);var n=e[0];var t=e[1];return(n+t)*3/4-t}function c(r,e,n){return(e+n)*3/4-n}function h(r){var e;var n=v(r);var t=n[0];var f=n[1];var i=new a(c(r,t,f));var u=0;var d=f>0?t-4:t;for(var h=0;h<d;h+=4){e=o[r.charCodeAt(h)]<<18|o[r.charCodeAt(h+1)]<<12|o[r.charCodeAt(h+2)]<<6|o[r.charCodeAt(h+3)];i[u++]=e>>16&255;i[u++]=e>>8&255;i[u++]=e&255}if(f===2){e=o[r.charCodeAt(h)]<<2|o[r.charCodeAt(h+1)]>>4;i[u++]=e&255}if(f===1){e=o[r.charCodeAt(h)]<<10|o[r.charCodeAt(h+1)]<<4|o[r.charCodeAt(h+2)]>>2;i[u++]=e>>8&255;i[u++]=e&255}return i}function s(r){return t[r>>18&63]+t[r>>12&63]+t[r>>6&63]+t[r&63]}function l(r,e,n){var t;var o=[];for(var a=e;a<n;a+=3){t=(r[a]<<16&16711680)+(r[a+1]<<8&65280)+(r[a+2]&255);o.push(s(t))}return o.join("")}function p(r){var e;var n=r.length;var o=n%3;var a=[];var f=16383;for(var i=0,u=n-o;i<u;i+=f){a.push(l(r,i,i+f>u?u:i+f))}if(o===1){e=r[n-1];a.push(t[e>>2]+t[e<<4&63]+"==")}else if(o===2){e=(r[n-2]<<8)+r[n-1];a.push(t[e>>10]+t[e>>4&63]+t[e<<2&63]+"=")}return a.join("")}},{}]},{},[])("/")});`

var inflateCode = `/** @license zlib.js 2012 - imaya [ https://github.com/imaya/zlib.js ] The MIT License */(function() {'use strict';var k=void 0,aa=this;function r(c,d){var a=c.split("."),b=aa;!(a[0]in b)&&b.execScript&&b.execScript("var "+a[0]);for(var e;a.length&&(e=a.shift());)!a.length&&d!==k?b[e]=d:b=b[e]?b[e]:b[e]={}};var t="undefined"!==typeof Uint8Array&&"undefined"!==typeof Uint16Array&&"undefined"!==typeof Uint32Array&&"undefined"!==typeof DataView;function u(c){var d=c.length,a=0,b=Number.POSITIVE_INFINITY,e,f,g,h,l,n,m,p,s,x;for(p=0;p<d;++p)c[p]>a&&(a=c[p]),c[p]<b&&(b=c[p]);e=1<<a;f=new (t?Uint32Array:Array)(e);g=1;h=0;for(l=2;g<=a;){for(p=0;p<d;++p)if(c[p]===g){n=0;m=h;for(s=0;s<g;++s)n=n<<1|m&1,m>>=1;x=g<<16|p;for(s=n;s<e;s+=l)f[s]=x;++h}++g;h<<=1;l<<=1}return[f,a,b]};function w(c,d){this.g=[];this.h=32768;this.c=this.f=this.d=this.k=0;this.input=t?new Uint8Array(c):c;this.l=!1;this.i=y;this.p=!1;if(d||!(d={}))d.index&&(this.d=d.index),d.bufferSize&&(this.h=d.bufferSize),d.bufferType&&(this.i=d.bufferType),d.resize&&(this.p=d.resize);switch(this.i){case A:this.a=32768;this.b=new (t?Uint8Array:Array)(32768+this.h+258);break;case y:this.a=0;this.b=new (t?Uint8Array:Array)(this.h);this.e=this.u;this.m=this.r;this.j=this.s;break;default:throw Error("invalid inflate mode");
}}var A=0,y=1;
w.prototype.t=function(){for(;!this.l;){var c=B(this,3);c&1&&(this.l=!0);c>>>=1;switch(c){case 0:var d=this.input,a=this.d,b=this.b,e=this.a,f=d.length,g=k,h=k,l=b.length,n=k;this.c=this.f=0;if(a+1>=f)throw Error("invalid uncompressed block header: LEN");g=d[a++]|d[a++]<<8;if(a+1>=f)throw Error("invalid uncompressed block header: NLEN");h=d[a++]|d[a++]<<8;if(g===~h)throw Error("invalid uncompressed block header: length verify");if(a+g>d.length)throw Error("input buffer is broken");switch(this.i){case A:for(;e+g>
b.length;){n=l-e;g-=n;if(t)b.set(d.subarray(a,a+n),e),e+=n,a+=n;else for(;n--;)b[e++]=d[a++];this.a=e;b=this.e();e=this.a}break;case y:for(;e+g>b.length;)b=this.e({o:2});break;default:throw Error("invalid inflate mode");}if(t)b.set(d.subarray(a,a+g),e),e+=g,a+=g;else for(;g--;)b[e++]=d[a++];this.d=a;this.a=e;this.b=b;break;case 1:this.j(ba,ca);break;case 2:for(var m=B(this,5)+257,p=B(this,5)+1,s=B(this,4)+4,x=new (t?Uint8Array:Array)(C.length),Q=k,R=k,S=k,v=k,M=k,F=k,z=k,q=k,T=k,q=0;q<s;++q)x[C[q]]=
B(this,3);if(!t){q=s;for(s=x.length;q<s;++q)x[C[q]]=0}Q=u(x);v=new (t?Uint8Array:Array)(m+p);q=0;for(T=m+p;q<T;)switch(M=D(this,Q),M){case 16:for(z=3+B(this,2);z--;)v[q++]=F;break;case 17:for(z=3+B(this,3);z--;)v[q++]=0;F=0;break;case 18:for(z=11+B(this,7);z--;)v[q++]=0;F=0;break;default:F=v[q++]=M}R=t?u(v.subarray(0,m)):u(v.slice(0,m));S=t?u(v.subarray(m)):u(v.slice(m));this.j(R,S);break;default:throw Error("unknown BTYPE: "+c);}}return this.m()};
var E=[16,17,18,0,8,7,9,6,10,5,11,4,12,3,13,2,14,1,15],C=t?new Uint16Array(E):E,G=[3,4,5,6,7,8,9,10,11,13,15,17,19,23,27,31,35,43,51,59,67,83,99,115,131,163,195,227,258,258,258],H=t?new Uint16Array(G):G,I=[0,0,0,0,0,0,0,0,1,1,1,1,2,2,2,2,3,3,3,3,4,4,4,4,5,5,5,5,0,0,0],J=t?new Uint8Array(I):I,K=[1,2,3,4,5,7,9,13,17,25,33,49,65,97,129,193,257,385,513,769,1025,1537,2049,3073,4097,6145,8193,12289,16385,24577],L=t?new Uint16Array(K):K,N=[0,0,0,0,1,1,2,2,3,3,4,4,5,5,6,6,7,7,8,8,9,9,10,10,11,11,12,12,13,
13],O=t?new Uint8Array(N):N,P=new (t?Uint8Array:Array)(288),U,da;U=0;for(da=P.length;U<da;++U)P[U]=143>=U?8:255>=U?9:279>=U?7:8;var ba=u(P),V=new (t?Uint8Array:Array)(30),W,ea;W=0;for(ea=V.length;W<ea;++W)V[W]=5;var ca=u(V);function B(c,d){for(var a=c.f,b=c.c,e=c.input,f=c.d,g=e.length,h;b<d;){if(f>=g)throw Error("input buffer is broken");a|=e[f++]<<b;b+=8}h=a&(1<<d)-1;c.f=a>>>d;c.c=b-d;c.d=f;return h}
function D(c,d){for(var a=c.f,b=c.c,e=c.input,f=c.d,g=e.length,h=d[0],l=d[1],n,m;b<l&&!(f>=g);)a|=e[f++]<<b,b+=8;n=h[a&(1<<l)-1];m=n>>>16;if(m>b)throw Error("invalid code length: "+m);c.f=a>>m;c.c=b-m;c.d=f;return n&65535}
w.prototype.j=function(c,d){var a=this.b,b=this.a;this.n=c;for(var e=a.length-258,f,g,h,l;256!==(f=D(this,c));)if(256>f)b>=e&&(this.a=b,a=this.e(),b=this.a),a[b++]=f;else{g=f-257;l=H[g];0<J[g]&&(l+=B(this,J[g]));f=D(this,d);h=L[f];0<O[f]&&(h+=B(this,O[f]));b>=e&&(this.a=b,a=this.e(),b=this.a);for(;l--;)a[b]=a[b++-h]}for(;8<=this.c;)this.c-=8,this.d--;this.a=b};
w.prototype.s=function(c,d){var a=this.b,b=this.a;this.n=c;for(var e=a.length,f,g,h,l;256!==(f=D(this,c));)if(256>f)b>=e&&(a=this.e(),e=a.length),a[b++]=f;else{g=f-257;l=H[g];0<J[g]&&(l+=B(this,J[g]));f=D(this,d);h=L[f];0<O[f]&&(h+=B(this,O[f]));b+l>e&&(a=this.e(),e=a.length);for(;l--;)a[b]=a[b++-h]}for(;8<=this.c;)this.c-=8,this.d--;this.a=b};
w.prototype.e=function(){var c=new (t?Uint8Array:Array)(this.a-32768),d=this.a-32768,a,b,e=this.b;if(t)c.set(e.subarray(32768,c.length));else{a=0;for(b=c.length;a<b;++a)c[a]=e[a+32768]}this.g.push(c);this.k+=c.length;if(t)e.set(e.subarray(d,d+32768));else for(a=0;32768>a;++a)e[a]=e[d+a];this.a=32768;return e};
w.prototype.u=function(c){var d,a=this.input.length/this.d+1|0,b,e,f,g=this.input,h=this.b;c&&("number"===typeof c.o&&(a=c.o),"number"===typeof c.q&&(a+=c.q));2>a?(b=(g.length-this.d)/this.n[2],f=258*(b/2)|0,e=f<h.length?h.length+f:h.length<<1):e=h.length*a;t?(d=new Uint8Array(e),d.set(h)):d=h;return this.b=d};
w.prototype.m=function(){var c=0,d=this.b,a=this.g,b,e=new (t?Uint8Array:Array)(this.k+(this.a-32768)),f,g,h,l;if(0===a.length)return t?this.b.subarray(32768,this.a):this.b.slice(32768,this.a);f=0;for(g=a.length;f<g;++f){b=a[f];h=0;for(l=b.length;h<l;++h)e[c++]=b[h]}f=32768;for(g=this.a;f<g;++f)e[c++]=d[f];this.g=[];return this.buffer=e};
w.prototype.r=function(){var c,d=this.a;t?this.p?(c=new Uint8Array(d),c.set(this.b.subarray(0,d))):c=this.b.subarray(0,d):(this.b.length>d&&(this.b.length=d),c=this.b);return this.buffer=c};r("Zlib.RawInflate",w);r("Zlib.RawInflate.prototype.decompress",w.prototype.t);var X={ADAPTIVE:y,BLOCK:A},Y,Z,$,fa;if(Object.keys)Y=Object.keys(X);else for(Z in Y=[],$=0,X)Y[$++]=Z;$=0;for(fa=Y.length;$<fa;++$)Z=Y[$],r("Zlib.RawInflate.BufferType."+Z,X[Z]);}).call(this);`

var goBootstrapCode = `(()=>{const e="undefined"!=typeof process;if(e){global.require=require,global.fs=require("fs");const e=require("crypto");global.crypto={getRandomValues(t){e.randomFillSync(t)}},global.performance={now(){const[e,t]=process.hrtime();return 1e3*e+t/1e6}};const t=require("util");global.TextEncoder=t.TextEncoder,global.TextDecoder=t.TextDecoder}else{if("undefined"!=typeof window)window.global=window;else{if("undefined"==typeof self)throw new Error("cannot export Go (neither window nor self is defined)");self.global=self}let e="";global.fs={constants:{O_WRONLY:-1,O_RDWR:-1,O_CREAT:-1,O_TRUNC:-1,O_APPEND:-1,O_EXCL:-1},writeSync(t,n){const i=(e+=s.decode(n)).lastIndexOf("\n");return-1!=i&&(console.log(e.substr(0,i)),e=e.substr(i+1)),n.length},openSync(e,t,s){const n=new Error("not implemented");throw n.code="ENOSYS",n}}}const t=new TextEncoder("utf-8"),s=new TextDecoder("utf-8");if(global.Go=class{constructor(){this.argv=["js"],this.env={},this.exit=(e=>{0!==e&&console.warn("exit code:",e)}),this._callbackTimeouts=new Map,this._nextCallbackTimeoutID=1;const e=()=>new DataView(this._inst.exports.mem.buffer),n=(t,s)=>{e().setUint32(t+0,s,!0),e().setUint32(t+4,Math.floor(s/4294967296),!0)},i=t=>{return e().getUint32(t+0,!0)+4294967296*e().getInt32(t+4,!0)},r=t=>{const s=e().getFloat64(t,!0);if(!isNaN(s))return s;const n=e().getUint32(t,!0);return this._values[n]},o=(t,s)=>{if("number"==typeof s)return isNaN(s)?(e().setUint32(t+4,2146959360,!0),void e().setUint32(t,0,!0)):void e().setFloat64(t,s,!0);switch(s){case void 0:return e().setUint32(t+4,2146959360,!0),void e().setUint32(t,1,!0);case null:return e().setUint32(t+4,2146959360,!0),void e().setUint32(t,2,!0);case!0:return e().setUint32(t+4,2146959360,!0),void e().setUint32(t,3,!0);case!1:return e().setUint32(t+4,2146959360,!0),void e().setUint32(t,4,!0)}let n=this._refs.get(s);void 0===n&&(n=this._values.length,this._values.push(s),this._refs.set(s,n));let i=0;switch(typeof s){case"string":i=1;break;case"symbol":i=2;break;case"function":i=3}e().setUint32(t+4,2146959360|i,!0),e().setUint32(t,n,!0)},a=e=>{const t=i(e+0),s=i(e+8);return new Uint8Array(this._inst.exports.mem.buffer,t,s)},l=e=>{const t=i(e+0),s=i(e+8),n=new Array(s);for(let e=0;e<s;e++)n[e]=r(t+8*e);return n},c=e=>{const t=i(e+0),n=i(e+8);return s.decode(new DataView(this._inst.exports.mem.buffer,t,n))},u=Date.now()-performance.now();this.importObject={go:{"runtime.wasmExit":t=>{const s=e().getInt32(t+8,!0);this.exited=!0,delete this._inst,delete this._values,delete this._refs,this.exit(s)},"runtime.wasmWrite":t=>{const s=i(t+8),n=i(t+16),r=e().getInt32(t+24,!0);fs.writeSync(s,new Uint8Array(this._inst.exports.mem.buffer,n,r))},"runtime.nanotime":e=>{n(e+8,1e6*(u+performance.now()))},"runtime.walltime":t=>{const s=(new Date).getTime();n(t+8,s/1e3),e().setInt32(t+16,s%1e3*1e6,!0)},"runtime.scheduleCallback":t=>{const s=this._nextCallbackTimeoutID;this._nextCallbackTimeoutID++,this._callbackTimeouts.set(s,setTimeout(()=>{this._resolveCallbackPromise()},i(t+8)+1)),e().setInt32(t+16,s,!0)},"runtime.clearScheduledCallback":t=>{const s=e().getInt32(t+8,!0);clearTimeout(this._callbackTimeouts.get(s)),this._callbackTimeouts.delete(s)},"runtime.getRandomData":e=>{crypto.getRandomValues(a(e+8))},"syscall/js.stringVal":e=>{o(e+24,c(e+8))},"syscall/js.valueGet":e=>{o(e+32,Reflect.get(r(e+8),c(e+16)))},"syscall/js.valueSet":e=>{Reflect.set(r(e+8),c(e+16),r(e+32))},"syscall/js.valueIndex":e=>{o(e+24,Reflect.get(r(e+8),i(e+16)))},"syscall/js.valueSetIndex":e=>{Reflect.set(r(e+8),i(e+16),r(e+24))},"syscall/js.valueCall":t=>{try{const s=r(t+8),n=Reflect.get(s,c(t+16)),i=l(t+32);o(t+56,Reflect.apply(n,s,i)),e().setUint8(t+64,1)}catch(s){o(t+56,s),e().setUint8(t+64,0)}},"syscall/js.valueInvoke":t=>{try{const s=r(t+8),n=l(t+16);o(t+40,Reflect.apply(s,void 0,n)),e().setUint8(t+48,1)}catch(s){o(t+40,s),e().setUint8(t+48,0)}},"syscall/js.valueNew":t=>{try{const s=r(t+8),n=l(t+16);o(t+40,Reflect.construct(s,n)),e().setUint8(t+48,1)}catch(s){o(t+40,s),e().setUint8(t+48,0)}},"syscall/js.valueLength":e=>{n(e+16,parseInt(r(e+8).length))},"syscall/js.valuePrepareString":e=>{const s=t.encode(String(r(e+8)));o(e+16,s),n(e+24,s.length)},"syscall/js.valueLoadString":e=>{const t=r(e+8);a(e+16).set(t)},"syscall/js.valueInstanceOf":t=>{e().setUint8(t+24,r(t+8)instanceof r(t+16))},debug:e=>{console.log(e)}}}}async run(e){this._inst=e,this._values=[NaN,void 0,null,!0,!1,global,this._inst.exports.mem,this],this._refs=new Map,this._callbackShutdown=!1,this.exited=!1;const s=new DataView(this._inst.exports.mem.buffer);let n=4096;const i=e=>{let i=n;return new Uint8Array(s.buffer,n,e.length+1).set(t.encode(e+"\0")),n+=e.length+(8-e.length%8),i},r=this.argv.length,o=[];this.argv.forEach(e=>{o.push(i(e))});const a=Object.keys(this.env).sort();o.push(a.length),a.forEach(e=>{o.push(i(e+"="+this.env[e]))});const l=n;for(o.forEach(e=>{s.setUint32(n,e,!0),s.setUint32(n+4,0,!0),n+=8});;){const e=new Promise(e=>{this._resolveCallbackPromise=(()=>{if(this.exited)throw new Error("bad callback: Go program has already exited");setTimeout(e,0)})});if(this._inst.exports.run(r,l),this.exited)break;await e}}static _makeCallbackHelper(e,t,s){return function(){t.push({id:e,args:arguments}),s._resolveCallbackPromise()}}static _makeEventCallbackHelper(e,t,s,n){return function(i){e&&i.preventDefault(),t&&i.stopPropagation(),s&&i.stopImmediatePropagation(),n(i)}}},e){process.argv.length<3&&(process.stderr.write("usage: go_js_wasm_exec [wasm binary] [arguments]\n"),process.exit(1));const e=new Go;e.argv=process.argv.slice(2),e.env=process.env,e.exit=process.exit,WebAssembly.instantiate(fs.readFileSync(process.argv[2]),e.importObject).then(t=>(process.on("exit",t=>{0!==t||e.exited||(e._callbackShutdown=!0,e._inst.exports.run())}),e.run(t.instance))).catch(e=>{throw e})}})();`

var opts struct {
	InputFile string `short:"i" long:"input" description:"JS input file"`
}

func main() {
	flags.Parse(&opts)
	if opts.InputFile == "" {
		fmt.Fprintf(os.Stderr, "Input file not specified\n")
		return
	}
	f, err := os.Open(opts.InputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %#v\n", err)
		return
	}
	input, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %#v\n", err)
		return
	}

	buf := &bytes.Buffer{}
	def, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating deflate object: %#v\n", err)
		return
	}
	n, err := def.Write(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compressing data: %#v\n", err)
		return
	}
	if len(input) != n {
		fmt.Fprintf(os.Stderr, "buffer size missmatch")
		return
	}
	err = def.Flush()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error flushing data: %#v\n", err)
		return
	}
	def.Close()
	b64str := base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Fprintf(os.Stdout,
		`/* Compressed with wasmpack [https://github.com/jaracil/wasmpack.git] */

%s
%s
%s
var go = new Go();
console.log("Before inflate", Date.now());
_bytes = new Zlib.RawInflate(base64js.toByteArray("%s")).decompress();
console.log("After inflate", Date.now());
window.onload = function(){
	WebAssembly.instantiate(_bytes , go.importObject).then(function(results){
		console.log("After instantiate", Date.now());
		go.run(results.instance);
	});
};
`,
		base64Code,
		inflateCode,
		goBootstrapCode,
		b64str,
	)

}
