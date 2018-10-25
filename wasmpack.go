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

var base64Code = `KGZ1bmN0aW9uKHIpe2lmKHR5cGVvZiBleHBvcnRzPT09Im9iamVjdCImJnR5cGVvZiBtb2R1bGUh
PT0idW5kZWZpbmVkIil7bW9kdWxlLmV4cG9ydHM9cigpfWVsc2UgaWYodHlwZW9mIGRlZmluZT09
PSJmdW5jdGlvbiImJmRlZmluZS5hbWQpe2RlZmluZShbXSxyKX1lbHNle3ZhciBlO2lmKHR5cGVv
ZiB3aW5kb3chPT0idW5kZWZpbmVkIil7ZT13aW5kb3d9ZWxzZSBpZih0eXBlb2YgZ2xvYmFsIT09
InVuZGVmaW5lZCIpe2U9Z2xvYmFsfWVsc2UgaWYodHlwZW9mIHNlbGYhPT0idW5kZWZpbmVkIil7
ZT1zZWxmfWVsc2V7ZT10aGlzfWUuYmFzZTY0anM9cigpfX0pKGZ1bmN0aW9uKCl7dmFyIHIsZSxu
O3JldHVybiBmdW5jdGlvbigpe2Z1bmN0aW9uIHIoZSxuLHQpe2Z1bmN0aW9uIG8oZixpKXtpZigh
bltmXSl7aWYoIWVbZl0pe3ZhciB1PSJmdW5jdGlvbiI9PXR5cGVvZiByZXF1aXJlJiZyZXF1aXJl
O2lmKCFpJiZ1KXJldHVybiB1KGYsITApO2lmKGEpcmV0dXJuIGEoZiwhMCk7dmFyIHY9bmV3IEVy
cm9yKCJDYW5ub3QgZmluZCBtb2R1bGUgJyIrZisiJyIpO3Rocm93IHYuY29kZT0iTU9EVUxFX05P
VF9GT1VORCIsdn12YXIgZD1uW2ZdPXtleHBvcnRzOnt9fTtlW2ZdWzBdLmNhbGwoZC5leHBvcnRz
LGZ1bmN0aW9uKHIpe3ZhciBuPWVbZl1bMV1bcl07cmV0dXJuIG8obnx8cil9LGQsZC5leHBvcnRz
LHIsZSxuLHQpfXJldHVybiBuW2ZdLmV4cG9ydHN9Zm9yKHZhciBhPSJmdW5jdGlvbiI9PXR5cGVv
ZiByZXF1aXJlJiZyZXF1aXJlLGY9MDtmPHQubGVuZ3RoO2YrKylvKHRbZl0pO3JldHVybiBvfXJl
dHVybiByfSgpKHsiLyI6W2Z1bmN0aW9uKHIsZSxuKXsidXNlIHN0cmljdCI7bi5ieXRlTGVuZ3Ro
PWQ7bi50b0J5dGVBcnJheT1oO24uZnJvbUJ5dGVBcnJheT1wO3ZhciB0PVtdO3ZhciBvPVtdO3Zh
ciBhPXR5cGVvZiBVaW50OEFycmF5IT09InVuZGVmaW5lZCI/VWludDhBcnJheTpBcnJheTt2YXIg
Zj0iQUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAx
MjM0NTY3ODkrLyI7Zm9yKHZhciBpPTAsdT1mLmxlbmd0aDtpPHU7KytpKXt0W2ldPWZbaV07b1tm
LmNoYXJDb2RlQXQoaSldPWl9b1siLSIuY2hhckNvZGVBdCgwKV09NjI7b1siXyIuY2hhckNvZGVB
dCgwKV09NjM7ZnVuY3Rpb24gdihyKXt2YXIgZT1yLmxlbmd0aDtpZihlJTQ+MCl7dGhyb3cgbmV3
IEVycm9yKCJJbnZhbGlkIHN0cmluZy4gTGVuZ3RoIG11c3QgYmUgYSBtdWx0aXBsZSBvZiA0Iil9
dmFyIG49ci5pbmRleE9mKCI9Iik7aWYobj09PS0xKW49ZTt2YXIgdD1uPT09ZT8wOjQtbiU0O3Jl
dHVybltuLHRdfWZ1bmN0aW9uIGQocil7dmFyIGU9dihyKTt2YXIgbj1lWzBdO3ZhciB0PWVbMV07
cmV0dXJuKG4rdCkqMy80LXR9ZnVuY3Rpb24gYyhyLGUsbil7cmV0dXJuKGUrbikqMy80LW59ZnVu
Y3Rpb24gaChyKXt2YXIgZTt2YXIgbj12KHIpO3ZhciB0PW5bMF07dmFyIGY9blsxXTt2YXIgaT1u
ZXcgYShjKHIsdCxmKSk7dmFyIHU9MDt2YXIgZD1mPjA/dC00OnQ7Zm9yKHZhciBoPTA7aDxkO2gr
PTQpe2U9b1tyLmNoYXJDb2RlQXQoaCldPDwxOHxvW3IuY2hhckNvZGVBdChoKzEpXTw8MTJ8b1ty
LmNoYXJDb2RlQXQoaCsyKV08PDZ8b1tyLmNoYXJDb2RlQXQoaCszKV07aVt1KytdPWU+PjE2JjI1
NTtpW3UrK109ZT4+OCYyNTU7aVt1KytdPWUmMjU1fWlmKGY9PT0yKXtlPW9bci5jaGFyQ29kZUF0
KGgpXTw8MnxvW3IuY2hhckNvZGVBdChoKzEpXT4+NDtpW3UrK109ZSYyNTV9aWYoZj09PTEpe2U9
b1tyLmNoYXJDb2RlQXQoaCldPDwxMHxvW3IuY2hhckNvZGVBdChoKzEpXTw8NHxvW3IuY2hhckNv
ZGVBdChoKzIpXT4+MjtpW3UrK109ZT4+OCYyNTU7aVt1KytdPWUmMjU1fXJldHVybiBpfWZ1bmN0
aW9uIHMocil7cmV0dXJuIHRbcj4+MTgmNjNdK3Rbcj4+MTImNjNdK3Rbcj4+NiY2M10rdFtyJjYz
XX1mdW5jdGlvbiBsKHIsZSxuKXt2YXIgdDt2YXIgbz1bXTtmb3IodmFyIGE9ZTthPG47YSs9Myl7
dD0oclthXTw8MTYmMTY3MTE2ODApKyhyW2ErMV08PDgmNjUyODApKyhyW2ErMl0mMjU1KTtvLnB1
c2gocyh0KSl9cmV0dXJuIG8uam9pbigiIil9ZnVuY3Rpb24gcChyKXt2YXIgZTt2YXIgbj1yLmxl
bmd0aDt2YXIgbz1uJTM7dmFyIGE9W107dmFyIGY9MTYzODM7Zm9yKHZhciBpPTAsdT1uLW87aTx1
O2krPWYpe2EucHVzaChsKHIsaSxpK2Y+dT91OmkrZikpfWlmKG89PT0xKXtlPXJbbi0xXTthLnB1
c2godFtlPj4yXSt0W2U8PDQmNjNdKyI9PSIpfWVsc2UgaWYobz09PTIpe2U9KHJbbi0yXTw8OCkr
cltuLTFdO2EucHVzaCh0W2U+PjEwXSt0W2U+PjQmNjNdK3RbZTw8MiY2M10rIj0iKX1yZXR1cm4g
YS5qb2luKCIiKX19LHt9XX0se30sW10pKCIvIil9KTsK`

var inflateCode = `LyoqIEBsaWNlbnNlIHpsaWIuanMgMjAxMiAtIGltYXlhIFsgaHR0cHM6Ly9naXRodWIuY29tL2lt
YXlhL3psaWIuanMgXSBUaGUgTUlUIExpY2Vuc2UgKi8oZnVuY3Rpb24oKSB7J3VzZSBzdHJpY3Qn
O3ZhciBrPXZvaWQgMCxhYT10aGlzO2Z1bmN0aW9uIHIoYyxkKXt2YXIgYT1jLnNwbGl0KCIuIiks
Yj1hYTshKGFbMF1pbiBiKSYmYi5leGVjU2NyaXB0JiZiLmV4ZWNTY3JpcHQoInZhciAiK2FbMF0p
O2Zvcih2YXIgZTthLmxlbmd0aCYmKGU9YS5zaGlmdCgpKTspIWEubGVuZ3RoJiZkIT09az9iW2Vd
PWQ6Yj1iW2VdP2JbZV06YltlXT17fX07dmFyIHQ9InVuZGVmaW5lZCIhPT10eXBlb2YgVWludDhB
cnJheSYmInVuZGVmaW5lZCIhPT10eXBlb2YgVWludDE2QXJyYXkmJiJ1bmRlZmluZWQiIT09dHlw
ZW9mIFVpbnQzMkFycmF5JiYidW5kZWZpbmVkIiE9PXR5cGVvZiBEYXRhVmlldztmdW5jdGlvbiB1
KGMpe3ZhciBkPWMubGVuZ3RoLGE9MCxiPU51bWJlci5QT1NJVElWRV9JTkZJTklUWSxlLGYsZyxo
LGwsbixtLHAscyx4O2ZvcihwPTA7cDxkOysrcCljW3BdPmEmJihhPWNbcF0pLGNbcF08YiYmKGI9
Y1twXSk7ZT0xPDxhO2Y9bmV3ICh0P1VpbnQzMkFycmF5OkFycmF5KShlKTtnPTE7aD0wO2Zvcihs
PTI7Zzw9YTspe2ZvcihwPTA7cDxkOysrcClpZihjW3BdPT09Zyl7bj0wO209aDtmb3Iocz0wO3M8
ZzsrK3Mpbj1uPDwxfG0mMSxtPj49MTt4PWc8PDE2fHA7Zm9yKHM9bjtzPGU7cys9bClmW3NdPXg7
KytofSsrZztoPDw9MTtsPDw9MX1yZXR1cm5bZixhLGJdfTtmdW5jdGlvbiB3KGMsZCl7dGhpcy5n
PVtdO3RoaXMuaD0zMjc2ODt0aGlzLmM9dGhpcy5mPXRoaXMuZD10aGlzLms9MDt0aGlzLmlucHV0
PXQ/bmV3IFVpbnQ4QXJyYXkoYyk6Yzt0aGlzLmw9ITE7dGhpcy5pPXk7dGhpcy5wPSExO2lmKGR8
fCEoZD17fSkpZC5pbmRleCYmKHRoaXMuZD1kLmluZGV4KSxkLmJ1ZmZlclNpemUmJih0aGlzLmg9
ZC5idWZmZXJTaXplKSxkLmJ1ZmZlclR5cGUmJih0aGlzLmk9ZC5idWZmZXJUeXBlKSxkLnJlc2l6
ZSYmKHRoaXMucD1kLnJlc2l6ZSk7c3dpdGNoKHRoaXMuaSl7Y2FzZSBBOnRoaXMuYT0zMjc2ODt0
aGlzLmI9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKDMyNzY4K3RoaXMuaCsyNTgpO2JyZWFrO2Nh
c2UgeTp0aGlzLmE9MDt0aGlzLmI9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKHRoaXMuaCk7dGhp
cy5lPXRoaXMudTt0aGlzLm09dGhpcy5yO3RoaXMuaj10aGlzLnM7YnJlYWs7ZGVmYXVsdDp0aHJv
dyBFcnJvcigiaW52YWxpZCBpbmZsYXRlIG1vZGUiKTsKfX12YXIgQT0wLHk9MTsKdy5wcm90b3R5
cGUudD1mdW5jdGlvbigpe2Zvcig7IXRoaXMubDspe3ZhciBjPUIodGhpcywzKTtjJjEmJih0aGlz
Lmw9ITApO2M+Pj49MTtzd2l0Y2goYyl7Y2FzZSAwOnZhciBkPXRoaXMuaW5wdXQsYT10aGlzLmQs
Yj10aGlzLmIsZT10aGlzLmEsZj1kLmxlbmd0aCxnPWssaD1rLGw9Yi5sZW5ndGgsbj1rO3RoaXMu
Yz10aGlzLmY9MDtpZihhKzE+PWYpdGhyb3cgRXJyb3IoImludmFsaWQgdW5jb21wcmVzc2VkIGJs
b2NrIGhlYWRlcjogTEVOIik7Zz1kW2ErK118ZFthKytdPDw4O2lmKGErMT49Zil0aHJvdyBFcnJv
cigiaW52YWxpZCB1bmNvbXByZXNzZWQgYmxvY2sgaGVhZGVyOiBOTEVOIik7aD1kW2ErK118ZFth
KytdPDw4O2lmKGc9PT1+aCl0aHJvdyBFcnJvcigiaW52YWxpZCB1bmNvbXByZXNzZWQgYmxvY2sg
aGVhZGVyOiBsZW5ndGggdmVyaWZ5Iik7aWYoYStnPmQubGVuZ3RoKXRocm93IEVycm9yKCJpbnB1
dCBidWZmZXIgaXMgYnJva2VuIik7c3dpdGNoKHRoaXMuaSl7Y2FzZSBBOmZvcig7ZStnPgpiLmxl
bmd0aDspe249bC1lO2ctPW47aWYodCliLnNldChkLnN1YmFycmF5KGEsYStuKSxlKSxlKz1uLGEr
PW47ZWxzZSBmb3IoO24tLTspYltlKytdPWRbYSsrXTt0aGlzLmE9ZTtiPXRoaXMuZSgpO2U9dGhp
cy5hfWJyZWFrO2Nhc2UgeTpmb3IoO2UrZz5iLmxlbmd0aDspYj10aGlzLmUoe286Mn0pO2JyZWFr
O2RlZmF1bHQ6dGhyb3cgRXJyb3IoImludmFsaWQgaW5mbGF0ZSBtb2RlIik7fWlmKHQpYi5zZXQo
ZC5zdWJhcnJheShhLGErZyksZSksZSs9ZyxhKz1nO2Vsc2UgZm9yKDtnLS07KWJbZSsrXT1kW2Er
K107dGhpcy5kPWE7dGhpcy5hPWU7dGhpcy5iPWI7YnJlYWs7Y2FzZSAxOnRoaXMuaihiYSxjYSk7
YnJlYWs7Y2FzZSAyOmZvcih2YXIgbT1CKHRoaXMsNSkrMjU3LHA9Qih0aGlzLDUpKzEscz1CKHRo
aXMsNCkrNCx4PW5ldyAodD9VaW50OEFycmF5OkFycmF5KShDLmxlbmd0aCksUT1rLFI9ayxTPWss
dj1rLE09ayxGPWssej1rLHE9ayxUPWsscT0wO3E8czsrK3EpeFtDW3FdXT0KQih0aGlzLDMpO2lm
KCF0KXtxPXM7Zm9yKHM9eC5sZW5ndGg7cTxzOysrcSl4W0NbcV1dPTB9UT11KHgpO3Y9bmV3ICh0
P1VpbnQ4QXJyYXk6QXJyYXkpKG0rcCk7cT0wO2ZvcihUPW0rcDtxPFQ7KXN3aXRjaChNPUQodGhp
cyxRKSxNKXtjYXNlIDE2OmZvcih6PTMrQih0aGlzLDIpO3otLTspdltxKytdPUY7YnJlYWs7Y2Fz
ZSAxNzpmb3Ioej0zK0IodGhpcywzKTt6LS07KXZbcSsrXT0wO0Y9MDticmVhaztjYXNlIDE4OmZv
cih6PTExK0IodGhpcyw3KTt6LS07KXZbcSsrXT0wO0Y9MDticmVhaztkZWZhdWx0OkY9dltxKytd
PU19Uj10P3Uodi5zdWJhcnJheSgwLG0pKTp1KHYuc2xpY2UoMCxtKSk7Uz10P3Uodi5zdWJhcnJh
eShtKSk6dSh2LnNsaWNlKG0pKTt0aGlzLmooUixTKTticmVhaztkZWZhdWx0OnRocm93IEVycm9y
KCJ1bmtub3duIEJUWVBFOiAiK2MpO319cmV0dXJuIHRoaXMubSgpfTsKdmFyIEU9WzE2LDE3LDE4
LDAsOCw3LDksNiwxMCw1LDExLDQsMTIsMywxMywyLDE0LDEsMTVdLEM9dD9uZXcgVWludDE2QXJy
YXkoRSk6RSxHPVszLDQsNSw2LDcsOCw5LDEwLDExLDEzLDE1LDE3LDE5LDIzLDI3LDMxLDM1LDQz
LDUxLDU5LDY3LDgzLDk5LDExNSwxMzEsMTYzLDE5NSwyMjcsMjU4LDI1OCwyNThdLEg9dD9uZXcg
VWludDE2QXJyYXkoRyk6RyxJPVswLDAsMCwwLDAsMCwwLDAsMSwxLDEsMSwyLDIsMiwyLDMsMywz
LDMsNCw0LDQsNCw1LDUsNSw1LDAsMCwwXSxKPXQ/bmV3IFVpbnQ4QXJyYXkoSSk6SSxLPVsxLDIs
Myw0LDUsNyw5LDEzLDE3LDI1LDMzLDQ5LDY1LDk3LDEyOSwxOTMsMjU3LDM4NSw1MTMsNzY5LDEw
MjUsMTUzNywyMDQ5LDMwNzMsNDA5Nyw2MTQ1LDgxOTMsMTIyODksMTYzODUsMjQ1NzddLEw9dD9u
ZXcgVWludDE2QXJyYXkoSyk6SyxOPVswLDAsMCwwLDEsMSwyLDIsMywzLDQsNCw1LDUsNiw2LDcs
Nyw4LDgsOSw5LDEwLDEwLDExLDExLDEyLDEyLDEzLAoxM10sTz10P25ldyBVaW50OEFycmF5KE4p
Ok4sUD1uZXcgKHQ/VWludDhBcnJheTpBcnJheSkoMjg4KSxVLGRhO1U9MDtmb3IoZGE9UC5sZW5n
dGg7VTxkYTsrK1UpUFtVXT0xNDM+PVU/ODoyNTU+PVU/OToyNzk+PVU/Nzo4O3ZhciBiYT11KFAp
LFY9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKDMwKSxXLGVhO1c9MDtmb3IoZWE9Vi5sZW5ndGg7
VzxlYTsrK1cpVltXXT01O3ZhciBjYT11KFYpO2Z1bmN0aW9uIEIoYyxkKXtmb3IodmFyIGE9Yy5m
LGI9Yy5jLGU9Yy5pbnB1dCxmPWMuZCxnPWUubGVuZ3RoLGg7YjxkOyl7aWYoZj49Zyl0aHJvdyBF
cnJvcigiaW5wdXQgYnVmZmVyIGlzIGJyb2tlbiIpO2F8PWVbZisrXTw8YjtiKz04fWg9YSYoMTw8
ZCktMTtjLmY9YT4+PmQ7Yy5jPWItZDtjLmQ9ZjtyZXR1cm4gaH0KZnVuY3Rpb24gRChjLGQpe2Zv
cih2YXIgYT1jLmYsYj1jLmMsZT1jLmlucHV0LGY9Yy5kLGc9ZS5sZW5ndGgsaD1kWzBdLGw9ZFsx
XSxuLG07YjxsJiYhKGY+PWcpOylhfD1lW2YrK108PGIsYis9ODtuPWhbYSYoMTw8bCktMV07bT1u
Pj4+MTY7aWYobT5iKXRocm93IEVycm9yKCJpbnZhbGlkIGNvZGUgbGVuZ3RoOiAiK20pO2MuZj1h
Pj5tO2MuYz1iLW07Yy5kPWY7cmV0dXJuIG4mNjU1MzV9CncucHJvdG90eXBlLmo9ZnVuY3Rpb24o
YyxkKXt2YXIgYT10aGlzLmIsYj10aGlzLmE7dGhpcy5uPWM7Zm9yKHZhciBlPWEubGVuZ3RoLTI1
OCxmLGcsaCxsOzI1NiE9PShmPUQodGhpcyxjKSk7KWlmKDI1Nj5mKWI+PWUmJih0aGlzLmE9Yixh
PXRoaXMuZSgpLGI9dGhpcy5hKSxhW2IrK109ZjtlbHNle2c9Zi0yNTc7bD1IW2ddOzA8SltnXSYm
KGwrPUIodGhpcyxKW2ddKSk7Zj1EKHRoaXMsZCk7aD1MW2ZdOzA8T1tmXSYmKGgrPUIodGhpcyxP
W2ZdKSk7Yj49ZSYmKHRoaXMuYT1iLGE9dGhpcy5lKCksYj10aGlzLmEpO2Zvcig7bC0tOylhW2Jd
PWFbYisrLWhdfWZvcig7ODw9dGhpcy5jOyl0aGlzLmMtPTgsdGhpcy5kLS07dGhpcy5hPWJ9Owp3
LnByb3RvdHlwZS5zPWZ1bmN0aW9uKGMsZCl7dmFyIGE9dGhpcy5iLGI9dGhpcy5hO3RoaXMubj1j
O2Zvcih2YXIgZT1hLmxlbmd0aCxmLGcsaCxsOzI1NiE9PShmPUQodGhpcyxjKSk7KWlmKDI1Nj5m
KWI+PWUmJihhPXRoaXMuZSgpLGU9YS5sZW5ndGgpLGFbYisrXT1mO2Vsc2V7Zz1mLTI1NztsPUhb
Z107MDxKW2ddJiYobCs9Qih0aGlzLEpbZ10pKTtmPUQodGhpcyxkKTtoPUxbZl07MDxPW2ZdJiYo
aCs9Qih0aGlzLE9bZl0pKTtiK2w+ZSYmKGE9dGhpcy5lKCksZT1hLmxlbmd0aCk7Zm9yKDtsLS07
KWFbYl09YVtiKystaF19Zm9yKDs4PD10aGlzLmM7KXRoaXMuYy09OCx0aGlzLmQtLTt0aGlzLmE9
Yn07CncucHJvdG90eXBlLmU9ZnVuY3Rpb24oKXt2YXIgYz1uZXcgKHQ/VWludDhBcnJheTpBcnJh
eSkodGhpcy5hLTMyNzY4KSxkPXRoaXMuYS0zMjc2OCxhLGIsZT10aGlzLmI7aWYodCljLnNldChl
LnN1YmFycmF5KDMyNzY4LGMubGVuZ3RoKSk7ZWxzZXthPTA7Zm9yKGI9Yy5sZW5ndGg7YTxiOysr
YSljW2FdPWVbYSszMjc2OF19dGhpcy5nLnB1c2goYyk7dGhpcy5rKz1jLmxlbmd0aDtpZih0KWUu
c2V0KGUuc3ViYXJyYXkoZCxkKzMyNzY4KSk7ZWxzZSBmb3IoYT0wOzMyNzY4PmE7KythKWVbYV09
ZVtkK2FdO3RoaXMuYT0zMjc2ODtyZXR1cm4gZX07CncucHJvdG90eXBlLnU9ZnVuY3Rpb24oYyl7
dmFyIGQsYT10aGlzLmlucHV0Lmxlbmd0aC90aGlzLmQrMXwwLGIsZSxmLGc9dGhpcy5pbnB1dCxo
PXRoaXMuYjtjJiYoIm51bWJlciI9PT10eXBlb2YgYy5vJiYoYT1jLm8pLCJudW1iZXIiPT09dHlw
ZW9mIGMucSYmKGErPWMucSkpOzI+YT8oYj0oZy5sZW5ndGgtdGhpcy5kKS90aGlzLm5bMl0sZj0y
NTgqKGIvMil8MCxlPWY8aC5sZW5ndGg/aC5sZW5ndGgrZjpoLmxlbmd0aDw8MSk6ZT1oLmxlbmd0
aCphO3Q/KGQ9bmV3IFVpbnQ4QXJyYXkoZSksZC5zZXQoaCkpOmQ9aDtyZXR1cm4gdGhpcy5iPWR9
Owp3LnByb3RvdHlwZS5tPWZ1bmN0aW9uKCl7dmFyIGM9MCxkPXRoaXMuYixhPXRoaXMuZyxiLGU9
bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKHRoaXMuaysodGhpcy5hLTMyNzY4KSksZixnLGgsbDtp
ZigwPT09YS5sZW5ndGgpcmV0dXJuIHQ/dGhpcy5iLnN1YmFycmF5KDMyNzY4LHRoaXMuYSk6dGhp
cy5iLnNsaWNlKDMyNzY4LHRoaXMuYSk7Zj0wO2ZvcihnPWEubGVuZ3RoO2Y8ZzsrK2Ype2I9YVtm
XTtoPTA7Zm9yKGw9Yi5sZW5ndGg7aDxsOysraCllW2MrK109YltoXX1mPTMyNzY4O2ZvcihnPXRo
aXMuYTtmPGc7KytmKWVbYysrXT1kW2ZdO3RoaXMuZz1bXTtyZXR1cm4gdGhpcy5idWZmZXI9ZX07
CncucHJvdG90eXBlLnI9ZnVuY3Rpb24oKXt2YXIgYyxkPXRoaXMuYTt0P3RoaXMucD8oYz1uZXcg
VWludDhBcnJheShkKSxjLnNldCh0aGlzLmIuc3ViYXJyYXkoMCxkKSkpOmM9dGhpcy5iLnN1YmFy
cmF5KDAsZCk6KHRoaXMuYi5sZW5ndGg+ZCYmKHRoaXMuYi5sZW5ndGg9ZCksYz10aGlzLmIpO3Jl
dHVybiB0aGlzLmJ1ZmZlcj1jfTtyKCJabGliLlJhd0luZmxhdGUiLHcpO3IoIlpsaWIuUmF3SW5m
bGF0ZS5wcm90b3R5cGUuZGVjb21wcmVzcyIsdy5wcm90b3R5cGUudCk7dmFyIFg9e0FEQVBUSVZF
OnksQkxPQ0s6QX0sWSxaLCQsZmE7aWYoT2JqZWN0LmtleXMpWT1PYmplY3Qua2V5cyhYKTtlbHNl
IGZvcihaIGluIFk9W10sJD0wLFgpWVskKytdPVo7JD0wO2ZvcihmYT1ZLmxlbmd0aDskPGZhOysr
JClaPVlbJF0scigiWmxpYi5SYXdJbmZsYXRlLkJ1ZmZlclR5cGUuIitaLFhbWl0pO30pLmNhbGwo
dGhpcyk7Cg==`

var goBootstrapCode = `KCgpPT57Y29uc3QgZT0idW5kZWZpbmVkIiE9dHlwZW9mIHByb2Nlc3M7aWYoZSl7Z2xvYmFsLnJl
cXVpcmU9cmVxdWlyZSxnbG9iYWwuZnM9cmVxdWlyZSgiZnMiKTtjb25zdCBlPXJlcXVpcmUoImNy
eXB0byIpO2dsb2JhbC5jcnlwdG89e2dldFJhbmRvbVZhbHVlcyh0KXtlLnJhbmRvbUZpbGxTeW5j
KHQpfX0sZ2xvYmFsLnBlcmZvcm1hbmNlPXtub3coKXtjb25zdFtlLHRdPXByb2Nlc3MuaHJ0aW1l
KCk7cmV0dXJuIDFlMyplK3QvMWU2fX07Y29uc3QgdD1yZXF1aXJlKCJ1dGlsIik7Z2xvYmFsLlRl
eHRFbmNvZGVyPXQuVGV4dEVuY29kZXIsZ2xvYmFsLlRleHREZWNvZGVyPXQuVGV4dERlY29kZXJ9
ZWxzZXtpZigidW5kZWZpbmVkIiE9dHlwZW9mIHdpbmRvdyl3aW5kb3cuZ2xvYmFsPXdpbmRvdztl
bHNle2lmKCJ1bmRlZmluZWQiPT10eXBlb2Ygc2VsZil0aHJvdyBuZXcgRXJyb3IoImNhbm5vdCBl
eHBvcnQgR28gKG5laXRoZXIgd2luZG93IG5vciBzZWxmIGlzIGRlZmluZWQpIik7c2VsZi5nbG9i
YWw9c2VsZn1sZXQgZT0iIjtnbG9iYWwuZnM9e2NvbnN0YW50czp7T19XUk9OTFk6LTEsT19SRFdS
Oi0xLE9fQ1JFQVQ6LTEsT19UUlVOQzotMSxPX0FQUEVORDotMSxPX0VYQ0w6LTF9LHdyaXRlU3lu
Yyh0LG4pe2NvbnN0IGk9KGUrPXMuZGVjb2RlKG4pKS5sYXN0SW5kZXhPZigiXG4iKTtyZXR1cm4t
MSE9aSYmKGNvbnNvbGUubG9nKGUuc3Vic3RyKDAsaSkpLGU9ZS5zdWJzdHIoaSsxKSksbi5sZW5n
dGh9LG9wZW5TeW5jKGUsdCxzKXtjb25zdCBuPW5ldyBFcnJvcigibm90IGltcGxlbWVudGVkIik7
dGhyb3cgbi5jb2RlPSJFTk9TWVMiLG59fX1jb25zdCB0PW5ldyBUZXh0RW5jb2RlcigidXRmLTgi
KSxzPW5ldyBUZXh0RGVjb2RlcigidXRmLTgiKTtpZihnbG9iYWwuR289Y2xhc3N7Y29uc3RydWN0
b3IoKXt0aGlzLmFyZ3Y9WyJqcyJdLHRoaXMuZW52PXt9LHRoaXMuZXhpdD0oZT0+ezAhPT1lJiZj
b25zb2xlLndhcm4oImV4aXQgY29kZToiLGUpfSksdGhpcy5fY2FsbGJhY2tUaW1lb3V0cz1uZXcg
TWFwLHRoaXMuX25leHRDYWxsYmFja1RpbWVvdXRJRD0xO2NvbnN0IGU9KCk9Pm5ldyBEYXRhVmll
dyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlciksbj0odCxzKT0+e2UoKS5zZXRVaW50MzIo
dCswLHMsITApLGUoKS5zZXRVaW50MzIodCs0LE1hdGguZmxvb3Iocy80Mjk0OTY3Mjk2KSwhMCl9
LGk9dD0+e3JldHVybiBlKCkuZ2V0VWludDMyKHQrMCwhMCkrNDI5NDk2NzI5NiplKCkuZ2V0SW50
MzIodCs0LCEwKX0scj10PT57Y29uc3Qgcz1lKCkuZ2V0RmxvYXQ2NCh0LCEwKTtpZighaXNOYU4o
cykpcmV0dXJuIHM7Y29uc3Qgbj1lKCkuZ2V0VWludDMyKHQsITApO3JldHVybiB0aGlzLl92YWx1
ZXNbbl19LG89KHQscyk9PntpZigibnVtYmVyIj09dHlwZW9mIHMpcmV0dXJuIGlzTmFOKHMpPyhl
KCkuc2V0VWludDMyKHQrNCwyMTQ2OTU5MzYwLCEwKSx2b2lkIGUoKS5zZXRVaW50MzIodCwwLCEw
KSk6dm9pZCBlKCkuc2V0RmxvYXQ2NCh0LHMsITApO3N3aXRjaChzKXtjYXNlIHZvaWQgMDpyZXR1
cm4gZSgpLnNldFVpbnQzMih0KzQsMjE0Njk1OTM2MCwhMCksdm9pZCBlKCkuc2V0VWludDMyKHQs
MSwhMCk7Y2FzZSBudWxsOnJldHVybiBlKCkuc2V0VWludDMyKHQrNCwyMTQ2OTU5MzYwLCEwKSx2
b2lkIGUoKS5zZXRVaW50MzIodCwyLCEwKTtjYXNlITA6cmV0dXJuIGUoKS5zZXRVaW50MzIodCs0
LDIxNDY5NTkzNjAsITApLHZvaWQgZSgpLnNldFVpbnQzMih0LDMsITApO2Nhc2UhMTpyZXR1cm4g
ZSgpLnNldFVpbnQzMih0KzQsMjE0Njk1OTM2MCwhMCksdm9pZCBlKCkuc2V0VWludDMyKHQsNCwh
MCl9bGV0IG49dGhpcy5fcmVmcy5nZXQocyk7dm9pZCAwPT09biYmKG49dGhpcy5fdmFsdWVzLmxl
bmd0aCx0aGlzLl92YWx1ZXMucHVzaChzKSx0aGlzLl9yZWZzLnNldChzLG4pKTtsZXQgaT0wO3N3
aXRjaCh0eXBlb2Ygcyl7Y2FzZSJzdHJpbmciOmk9MTticmVhaztjYXNlInN5bWJvbCI6aT0yO2Jy
ZWFrO2Nhc2UiZnVuY3Rpb24iOmk9M31lKCkuc2V0VWludDMyKHQrNCwyMTQ2OTU5MzYwfGksITAp
LGUoKS5zZXRVaW50MzIodCxuLCEwKX0sYT1lPT57Y29uc3QgdD1pKGUrMCkscz1pKGUrOCk7cmV0
dXJuIG5ldyBVaW50OEFycmF5KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyLHQscyl9LGw9
ZT0+e2NvbnN0IHQ9aShlKzApLHM9aShlKzgpLG49bmV3IEFycmF5KHMpO2ZvcihsZXQgZT0wO2U8
cztlKyspbltlXT1yKHQrOCplKTtyZXR1cm4gbn0sYz1lPT57Y29uc3QgdD1pKGUrMCksbj1pKGUr
OCk7cmV0dXJuIHMuZGVjb2RlKG5ldyBEYXRhVmlldyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1
ZmZlcix0LG4pKX0sdT1EYXRlLm5vdygpLXBlcmZvcm1hbmNlLm5vdygpO3RoaXMuaW1wb3J0T2Jq
ZWN0PXtnbzp7InJ1bnRpbWUud2FzbUV4aXQiOnQ9Pntjb25zdCBzPWUoKS5nZXRJbnQzMih0Kzgs
ITApO3RoaXMuZXhpdGVkPSEwLGRlbGV0ZSB0aGlzLl9pbnN0LGRlbGV0ZSB0aGlzLl92YWx1ZXMs
ZGVsZXRlIHRoaXMuX3JlZnMsdGhpcy5leGl0KHMpfSwicnVudGltZS53YXNtV3JpdGUiOnQ9Pntj
b25zdCBzPWkodCs4KSxuPWkodCsxNikscj1lKCkuZ2V0SW50MzIodCsyNCwhMCk7ZnMud3JpdGVT
eW5jKHMsbmV3IFVpbnQ4QXJyYXkodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIsbixyKSl9
LCJydW50aW1lLm5hbm90aW1lIjplPT57bihlKzgsMWU2Kih1K3BlcmZvcm1hbmNlLm5vdygpKSl9
LCJydW50aW1lLndhbGx0aW1lIjp0PT57Y29uc3Qgcz0obmV3IERhdGUpLmdldFRpbWUoKTtuKHQr
OCxzLzFlMyksZSgpLnNldEludDMyKHQrMTYscyUxZTMqMWU2LCEwKX0sInJ1bnRpbWUuc2NoZWR1
bGVDYWxsYmFjayI6dD0+e2NvbnN0IHM9dGhpcy5fbmV4dENhbGxiYWNrVGltZW91dElEO3RoaXMu
X25leHRDYWxsYmFja1RpbWVvdXRJRCsrLHRoaXMuX2NhbGxiYWNrVGltZW91dHMuc2V0KHMsc2V0
VGltZW91dCgoKT0+e3RoaXMuX3Jlc29sdmVDYWxsYmFja1Byb21pc2UoKX0saSh0KzgpKzEpKSxl
KCkuc2V0SW50MzIodCsxNixzLCEwKX0sInJ1bnRpbWUuY2xlYXJTY2hlZHVsZWRDYWxsYmFjayI6
dD0+e2NvbnN0IHM9ZSgpLmdldEludDMyKHQrOCwhMCk7Y2xlYXJUaW1lb3V0KHRoaXMuX2NhbGxi
YWNrVGltZW91dHMuZ2V0KHMpKSx0aGlzLl9jYWxsYmFja1RpbWVvdXRzLmRlbGV0ZShzKX0sInJ1
bnRpbWUuZ2V0UmFuZG9tRGF0YSI6ZT0+e2NyeXB0by5nZXRSYW5kb21WYWx1ZXMoYShlKzgpKX0s
InN5c2NhbGwvanMuc3RyaW5nVmFsIjplPT57byhlKzI0LGMoZSs4KSl9LCJzeXNjYWxsL2pzLnZh
bHVlR2V0IjplPT57byhlKzMyLFJlZmxlY3QuZ2V0KHIoZSs4KSxjKGUrMTYpKSl9LCJzeXNjYWxs
L2pzLnZhbHVlU2V0IjplPT57UmVmbGVjdC5zZXQocihlKzgpLGMoZSsxNikscihlKzMyKSl9LCJz
eXNjYWxsL2pzLnZhbHVlSW5kZXgiOmU9PntvKGUrMjQsUmVmbGVjdC5nZXQocihlKzgpLGkoZSsx
NikpKX0sInN5c2NhbGwvanMudmFsdWVTZXRJbmRleCI6ZT0+e1JlZmxlY3Quc2V0KHIoZSs4KSxp
KGUrMTYpLHIoZSsyNCkpfSwic3lzY2FsbC9qcy52YWx1ZUNhbGwiOnQ9Pnt0cnl7Y29uc3Qgcz1y
KHQrOCksbj1SZWZsZWN0LmdldChzLGModCsxNikpLGk9bCh0KzMyKTtvKHQrNTYsUmVmbGVjdC5h
cHBseShuLHMsaSkpLGUoKS5zZXRVaW50OCh0KzY0LDEpfWNhdGNoKHMpe28odCs1NixzKSxlKCku
c2V0VWludDgodCs2NCwwKX19LCJzeXNjYWxsL2pzLnZhbHVlSW52b2tlIjp0PT57dHJ5e2NvbnN0
IHM9cih0KzgpLG49bCh0KzE2KTtvKHQrNDAsUmVmbGVjdC5hcHBseShzLHZvaWQgMCxuKSksZSgp
LnNldFVpbnQ4KHQrNDgsMSl9Y2F0Y2gocyl7byh0KzQwLHMpLGUoKS5zZXRVaW50OCh0KzQ4LDAp
fX0sInN5c2NhbGwvanMudmFsdWVOZXciOnQ9Pnt0cnl7Y29uc3Qgcz1yKHQrOCksbj1sKHQrMTYp
O28odCs0MCxSZWZsZWN0LmNvbnN0cnVjdChzLG4pKSxlKCkuc2V0VWludDgodCs0OCwxKX1jYXRj
aChzKXtvKHQrNDAscyksZSgpLnNldFVpbnQ4KHQrNDgsMCl9fSwic3lzY2FsbC9qcy52YWx1ZUxl
bmd0aCI6ZT0+e24oZSsxNixwYXJzZUludChyKGUrOCkubGVuZ3RoKSl9LCJzeXNjYWxsL2pzLnZh
bHVlUHJlcGFyZVN0cmluZyI6ZT0+e2NvbnN0IHM9dC5lbmNvZGUoU3RyaW5nKHIoZSs4KSkpO28o
ZSsxNixzKSxuKGUrMjQscy5sZW5ndGgpfSwic3lzY2FsbC9qcy52YWx1ZUxvYWRTdHJpbmciOmU9
Pntjb25zdCB0PXIoZSs4KTthKGUrMTYpLnNldCh0KX0sInN5c2NhbGwvanMudmFsdWVJbnN0YW5j
ZU9mIjp0PT57ZSgpLnNldFVpbnQ4KHQrMjQscih0KzgpaW5zdGFuY2VvZiByKHQrMTYpKX0sZGVi
dWc6ZT0+e2NvbnNvbGUubG9nKGUpfX19fWFzeW5jIHJ1bihlKXt0aGlzLl9pbnN0PWUsdGhpcy5f
dmFsdWVzPVtOYU4sdm9pZCAwLG51bGwsITAsITEsZ2xvYmFsLHRoaXMuX2luc3QuZXhwb3J0cy5t
ZW0sdGhpc10sdGhpcy5fcmVmcz1uZXcgTWFwLHRoaXMuX2NhbGxiYWNrU2h1dGRvd249ITEsdGhp
cy5leGl0ZWQ9ITE7Y29uc3Qgcz1uZXcgRGF0YVZpZXcodGhpcy5faW5zdC5leHBvcnRzLm1lbS5i
dWZmZXIpO2xldCBuPTQwOTY7Y29uc3QgaT1lPT57bGV0IGk9bjtyZXR1cm4gbmV3IFVpbnQ4QXJy
YXkocy5idWZmZXIsbixlLmxlbmd0aCsxKS5zZXQodC5lbmNvZGUoZSsiXDAiKSksbis9ZS5sZW5n
dGgrKDgtZS5sZW5ndGglOCksaX0scj10aGlzLmFyZ3YubGVuZ3RoLG89W107dGhpcy5hcmd2LmZv
ckVhY2goZT0+e28ucHVzaChpKGUpKX0pO2NvbnN0IGE9T2JqZWN0LmtleXModGhpcy5lbnYpLnNv
cnQoKTtvLnB1c2goYS5sZW5ndGgpLGEuZm9yRWFjaChlPT57by5wdXNoKGkoYCR7ZX09JHt0aGlz
LmVudltlXX1gKSl9KTtjb25zdCBsPW47Zm9yKG8uZm9yRWFjaChlPT57cy5zZXRVaW50MzIobixl
LCEwKSxzLnNldFVpbnQzMihuKzQsMCwhMCksbis9OH0pOzspe2NvbnN0IGU9bmV3IFByb21pc2Uo
ZT0+e3RoaXMuX3Jlc29sdmVDYWxsYmFja1Byb21pc2U9KCgpPT57aWYodGhpcy5leGl0ZWQpdGhy
b3cgbmV3IEVycm9yKCJiYWQgY2FsbGJhY2s6IEdvIHByb2dyYW0gaGFzIGFscmVhZHkgZXhpdGVk
Iik7c2V0VGltZW91dChlLDApfSl9KTtpZih0aGlzLl9pbnN0LmV4cG9ydHMucnVuKHIsbCksdGhp
cy5leGl0ZWQpYnJlYWs7YXdhaXQgZX19c3RhdGljIF9tYWtlQ2FsbGJhY2tIZWxwZXIoZSx0LHMp
e3JldHVybiBmdW5jdGlvbigpe3QucHVzaCh7aWQ6ZSxhcmdzOmFyZ3VtZW50c30pLHMuX3Jlc29s
dmVDYWxsYmFja1Byb21pc2UoKX19c3RhdGljIF9tYWtlRXZlbnRDYWxsYmFja0hlbHBlcihlLHQs
cyxuKXtyZXR1cm4gZnVuY3Rpb24oaSl7ZSYmaS5wcmV2ZW50RGVmYXVsdCgpLHQmJmkuc3RvcFBy
b3BhZ2F0aW9uKCkscyYmaS5zdG9wSW1tZWRpYXRlUHJvcGFnYXRpb24oKSxuKGkpfX19LGUpe3By
b2Nlc3MuYXJndi5sZW5ndGg8MyYmKHByb2Nlc3Muc3RkZXJyLndyaXRlKCJ1c2FnZTogZ29fanNf
d2FzbV9leGVjIFt3YXNtIGJpbmFyeV0gW2FyZ3VtZW50c11cbiIpLHByb2Nlc3MuZXhpdCgxKSk7
Y29uc3QgZT1uZXcgR287ZS5hcmd2PXByb2Nlc3MuYXJndi5zbGljZSgyKSxlLmVudj1wcm9jZXNz
LmVudixlLmV4aXQ9cHJvY2Vzcy5leGl0LFdlYkFzc2VtYmx5Lmluc3RhbnRpYXRlKGZzLnJlYWRG
aWxlU3luYyhwcm9jZXNzLmFyZ3ZbMl0pLGUuaW1wb3J0T2JqZWN0KS50aGVuKHQ9Pihwcm9jZXNz
Lm9uKCJleGl0Iix0PT57MCE9PXR8fGUuZXhpdGVkfHwoZS5fY2FsbGJhY2tTaHV0ZG93bj0hMCxl
Ll9pbnN0LmV4cG9ydHMucnVuKCkpfSksZS5ydW4odC5pbnN0YW5jZSkpKS5jYXRjaChlPT57dGhy
b3cgZX0pfX0pKCk7Cg==`

var opts struct {
	InputFile string `short:"i" long:"input" description:"WASM input file"`
}

func genCode(s string) string {
	bin, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(bin)
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
window.onload = function(){
	WebAssembly.instantiate(new Zlib.RawInflate(base64js.toByteArray("%s")).decompress() , go.importObject).then(function(results){
		go.run(results.instance);
	});
};
`,
		genCode(base64Code),
		genCode(inflateCode),
		genCode(goBootstrapCode),
		b64str,
	)

}
