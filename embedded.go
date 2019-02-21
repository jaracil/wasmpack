package main

var base64Code = "KGZ1bmN0aW9uKHIpe2lmKHR5cGVvZiBleHBvcnRzPT09Im9iamVjdCImJnR5cGVvZiBtb2R1bGUhPT0idW5kZWZpbmVkIil7bW9kdWxlLmV4cG9ydHM9cigpfWVsc2UgaWYodHlwZW9mIGRlZmluZT09PSJmdW5jdGlvbiImJmRlZmluZS5hbWQpe2RlZmluZShbXSxyKX1lbHNle3ZhciBlO2lmKHR5cGVvZiB3aW5kb3chPT0idW5kZWZpbmVkIil7ZT13aW5kb3d9ZWxzZSBpZih0eXBlb2YgZ2xvYmFsIT09InVuZGVmaW5lZCIpe2U9Z2xvYmFsfWVsc2UgaWYodHlwZW9mIHNlbGYhPT0idW5kZWZpbmVkIil7ZT1zZWxmfWVsc2V7ZT10aGlzfWUuYmFzZTY0anM9cigpfX0pKGZ1bmN0aW9uKCl7dmFyIHIsZSxuO3JldHVybiBmdW5jdGlvbigpe2Z1bmN0aW9uIHIoZSxuLHQpe2Z1bmN0aW9uIG8oZixpKXtpZighbltmXSl7aWYoIWVbZl0pe3ZhciB1PSJmdW5jdGlvbiI9PXR5cGVvZiByZXF1aXJlJiZyZXF1aXJlO2lmKCFpJiZ1KXJldHVybiB1KGYsITApO2lmKGEpcmV0dXJuIGEoZiwhMCk7dmFyIHY9bmV3IEVycm9yKCJDYW5ub3QgZmluZCBtb2R1bGUgJyIrZisiJyIpO3Rocm93IHYuY29kZT0iTU9EVUxFX05PVF9GT1VORCIsdn12YXIgZD1uW2ZdPXtleHBvcnRzOnt9fTtlW2ZdWzBdLmNhbGwoZC5leHBvcnRzLGZ1bmN0aW9uKHIpe3ZhciBuPWVbZl1bMV1bcl07cmV0dXJuIG8obnx8cil9LGQsZC5leHBvcnRzLHIsZSxuLHQpfXJldHVybiBuW2ZdLmV4cG9ydHN9Zm9yKHZhciBhPSJmdW5jdGlvbiI9PXR5cGVvZiByZXF1aXJlJiZyZXF1aXJlLGY9MDtmPHQubGVuZ3RoO2YrKylvKHRbZl0pO3JldHVybiBvfXJldHVybiByfSgpKHsiLyI6W2Z1bmN0aW9uKHIsZSxuKXsidXNlIHN0cmljdCI7bi5ieXRlTGVuZ3RoPWQ7bi50b0J5dGVBcnJheT1oO24uZnJvbUJ5dGVBcnJheT1wO3ZhciB0PVtdO3ZhciBvPVtdO3ZhciBhPXR5cGVvZiBVaW50OEFycmF5IT09InVuZGVmaW5lZCI/VWludDhBcnJheTpBcnJheTt2YXIgZj0iQUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODkrLyI7Zm9yKHZhciBpPTAsdT1mLmxlbmd0aDtpPHU7KytpKXt0W2ldPWZbaV07b1tmLmNoYXJDb2RlQXQoaSldPWl9b1siLSIuY2hhckNvZGVBdCgwKV09NjI7b1siXyIuY2hhckNvZGVBdCgwKV09NjM7ZnVuY3Rpb24gdihyKXt2YXIgZT1yLmxlbmd0aDtpZihlJTQ+MCl7dGhyb3cgbmV3IEVycm9yKCJJbnZhbGlkIHN0cmluZy4gTGVuZ3RoIG11c3QgYmUgYSBtdWx0aXBsZSBvZiA0Iil9dmFyIG49ci5pbmRleE9mKCI9Iik7aWYobj09PS0xKW49ZTt2YXIgdD1uPT09ZT8wOjQtbiU0O3JldHVybltuLHRdfWZ1bmN0aW9uIGQocil7dmFyIGU9dihyKTt2YXIgbj1lWzBdO3ZhciB0PWVbMV07cmV0dXJuKG4rdCkqMy80LXR9ZnVuY3Rpb24gYyhyLGUsbil7cmV0dXJuKGUrbikqMy80LW59ZnVuY3Rpb24gaChyKXt2YXIgZTt2YXIgbj12KHIpO3ZhciB0PW5bMF07dmFyIGY9blsxXTt2YXIgaT1uZXcgYShjKHIsdCxmKSk7dmFyIHU9MDt2YXIgZD1mPjA/dC00OnQ7Zm9yKHZhciBoPTA7aDxkO2grPTQpe2U9b1tyLmNoYXJDb2RlQXQoaCldPDwxOHxvW3IuY2hhckNvZGVBdChoKzEpXTw8MTJ8b1tyLmNoYXJDb2RlQXQoaCsyKV08PDZ8b1tyLmNoYXJDb2RlQXQoaCszKV07aVt1KytdPWU+PjE2JjI1NTtpW3UrK109ZT4+OCYyNTU7aVt1KytdPWUmMjU1fWlmKGY9PT0yKXtlPW9bci5jaGFyQ29kZUF0KGgpXTw8MnxvW3IuY2hhckNvZGVBdChoKzEpXT4+NDtpW3UrK109ZSYyNTV9aWYoZj09PTEpe2U9b1tyLmNoYXJDb2RlQXQoaCldPDwxMHxvW3IuY2hhckNvZGVBdChoKzEpXTw8NHxvW3IuY2hhckNvZGVBdChoKzIpXT4+MjtpW3UrK109ZT4+OCYyNTU7aVt1KytdPWUmMjU1fXJldHVybiBpfWZ1bmN0aW9uIHMocil7cmV0dXJuIHRbcj4+MTgmNjNdK3Rbcj4+MTImNjNdK3Rbcj4+NiY2M10rdFtyJjYzXX1mdW5jdGlvbiBsKHIsZSxuKXt2YXIgdDt2YXIgbz1bXTtmb3IodmFyIGE9ZTthPG47YSs9Myl7dD0oclthXTw8MTYmMTY3MTE2ODApKyhyW2ErMV08PDgmNjUyODApKyhyW2ErMl0mMjU1KTtvLnB1c2gocyh0KSl9cmV0dXJuIG8uam9pbigiIil9ZnVuY3Rpb24gcChyKXt2YXIgZTt2YXIgbj1yLmxlbmd0aDt2YXIgbz1uJTM7dmFyIGE9W107dmFyIGY9MTYzODM7Zm9yKHZhciBpPTAsdT1uLW87aTx1O2krPWYpe2EucHVzaChsKHIsaSxpK2Y+dT91OmkrZikpfWlmKG89PT0xKXtlPXJbbi0xXTthLnB1c2godFtlPj4yXSt0W2U8PDQmNjNdKyI9PSIpfWVsc2UgaWYobz09PTIpe2U9KHJbbi0yXTw8OCkrcltuLTFdO2EucHVzaCh0W2U+PjEwXSt0W2U+PjQmNjNdK3RbZTw8MiY2M10rIj0iKX1yZXR1cm4gYS5qb2luKCIiKX19LHt9XX0se30sW10pKCIvIil9KTsK"

var inflateCode = "LyoqIEBsaWNlbnNlIHpsaWIuanMgMjAxMiAtIGltYXlhIFsgaHR0cHM6Ly9naXRodWIuY29tL2ltYXlhL3psaWIuanMgXSBUaGUgTUlUIExpY2Vuc2UgKi8oZnVuY3Rpb24oKSB7J3VzZSBzdHJpY3QnO3ZhciBrPXZvaWQgMCxhYT10aGlzO2Z1bmN0aW9uIHIoYyxkKXt2YXIgYT1jLnNwbGl0KCIuIiksYj1hYTshKGFbMF1pbiBiKSYmYi5leGVjU2NyaXB0JiZiLmV4ZWNTY3JpcHQoInZhciAiK2FbMF0pO2Zvcih2YXIgZTthLmxlbmd0aCYmKGU9YS5zaGlmdCgpKTspIWEubGVuZ3RoJiZkIT09az9iW2VdPWQ6Yj1iW2VdP2JbZV06YltlXT17fX07dmFyIHQ9InVuZGVmaW5lZCIhPT10eXBlb2YgVWludDhBcnJheSYmInVuZGVmaW5lZCIhPT10eXBlb2YgVWludDE2QXJyYXkmJiJ1bmRlZmluZWQiIT09dHlwZW9mIFVpbnQzMkFycmF5JiYidW5kZWZpbmVkIiE9PXR5cGVvZiBEYXRhVmlldztmdW5jdGlvbiB1KGMpe3ZhciBkPWMubGVuZ3RoLGE9MCxiPU51bWJlci5QT1NJVElWRV9JTkZJTklUWSxlLGYsZyxoLGwsbixtLHAscyx4O2ZvcihwPTA7cDxkOysrcCljW3BdPmEmJihhPWNbcF0pLGNbcF08YiYmKGI9Y1twXSk7ZT0xPDxhO2Y9bmV3ICh0P1VpbnQzMkFycmF5OkFycmF5KShlKTtnPTE7aD0wO2ZvcihsPTI7Zzw9YTspe2ZvcihwPTA7cDxkOysrcClpZihjW3BdPT09Zyl7bj0wO209aDtmb3Iocz0wO3M8ZzsrK3Mpbj1uPDwxfG0mMSxtPj49MTt4PWc8PDE2fHA7Zm9yKHM9bjtzPGU7cys9bClmW3NdPXg7KytofSsrZztoPDw9MTtsPDw9MX1yZXR1cm5bZixhLGJdfTtmdW5jdGlvbiB3KGMsZCl7dGhpcy5nPVtdO3RoaXMuaD0zMjc2ODt0aGlzLmM9dGhpcy5mPXRoaXMuZD10aGlzLms9MDt0aGlzLmlucHV0PXQ/bmV3IFVpbnQ4QXJyYXkoYyk6Yzt0aGlzLmw9ITE7dGhpcy5pPXk7dGhpcy5wPSExO2lmKGR8fCEoZD17fSkpZC5pbmRleCYmKHRoaXMuZD1kLmluZGV4KSxkLmJ1ZmZlclNpemUmJih0aGlzLmg9ZC5idWZmZXJTaXplKSxkLmJ1ZmZlclR5cGUmJih0aGlzLmk9ZC5idWZmZXJUeXBlKSxkLnJlc2l6ZSYmKHRoaXMucD1kLnJlc2l6ZSk7c3dpdGNoKHRoaXMuaSl7Y2FzZSBBOnRoaXMuYT0zMjc2ODt0aGlzLmI9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKDMyNzY4K3RoaXMuaCsyNTgpO2JyZWFrO2Nhc2UgeTp0aGlzLmE9MDt0aGlzLmI9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKHRoaXMuaCk7dGhpcy5lPXRoaXMudTt0aGlzLm09dGhpcy5yO3RoaXMuaj10aGlzLnM7YnJlYWs7ZGVmYXVsdDp0aHJvdyBFcnJvcigiaW52YWxpZCBpbmZsYXRlIG1vZGUiKTsKfX12YXIgQT0wLHk9MTsKdy5wcm90b3R5cGUudD1mdW5jdGlvbigpe2Zvcig7IXRoaXMubDspe3ZhciBjPUIodGhpcywzKTtjJjEmJih0aGlzLmw9ITApO2M+Pj49MTtzd2l0Y2goYyl7Y2FzZSAwOnZhciBkPXRoaXMuaW5wdXQsYT10aGlzLmQsYj10aGlzLmIsZT10aGlzLmEsZj1kLmxlbmd0aCxnPWssaD1rLGw9Yi5sZW5ndGgsbj1rO3RoaXMuYz10aGlzLmY9MDtpZihhKzE+PWYpdGhyb3cgRXJyb3IoImludmFsaWQgdW5jb21wcmVzc2VkIGJsb2NrIGhlYWRlcjogTEVOIik7Zz1kW2ErK118ZFthKytdPDw4O2lmKGErMT49Zil0aHJvdyBFcnJvcigiaW52YWxpZCB1bmNvbXByZXNzZWQgYmxvY2sgaGVhZGVyOiBOTEVOIik7aD1kW2ErK118ZFthKytdPDw4O2lmKGc9PT1+aCl0aHJvdyBFcnJvcigiaW52YWxpZCB1bmNvbXByZXNzZWQgYmxvY2sgaGVhZGVyOiBsZW5ndGggdmVyaWZ5Iik7aWYoYStnPmQubGVuZ3RoKXRocm93IEVycm9yKCJpbnB1dCBidWZmZXIgaXMgYnJva2VuIik7c3dpdGNoKHRoaXMuaSl7Y2FzZSBBOmZvcig7ZStnPgpiLmxlbmd0aDspe249bC1lO2ctPW47aWYodCliLnNldChkLnN1YmFycmF5KGEsYStuKSxlKSxlKz1uLGErPW47ZWxzZSBmb3IoO24tLTspYltlKytdPWRbYSsrXTt0aGlzLmE9ZTtiPXRoaXMuZSgpO2U9dGhpcy5hfWJyZWFrO2Nhc2UgeTpmb3IoO2UrZz5iLmxlbmd0aDspYj10aGlzLmUoe286Mn0pO2JyZWFrO2RlZmF1bHQ6dGhyb3cgRXJyb3IoImludmFsaWQgaW5mbGF0ZSBtb2RlIik7fWlmKHQpYi5zZXQoZC5zdWJhcnJheShhLGErZyksZSksZSs9ZyxhKz1nO2Vsc2UgZm9yKDtnLS07KWJbZSsrXT1kW2ErK107dGhpcy5kPWE7dGhpcy5hPWU7dGhpcy5iPWI7YnJlYWs7Y2FzZSAxOnRoaXMuaihiYSxjYSk7YnJlYWs7Y2FzZSAyOmZvcih2YXIgbT1CKHRoaXMsNSkrMjU3LHA9Qih0aGlzLDUpKzEscz1CKHRoaXMsNCkrNCx4PW5ldyAodD9VaW50OEFycmF5OkFycmF5KShDLmxlbmd0aCksUT1rLFI9ayxTPWssdj1rLE09ayxGPWssej1rLHE9ayxUPWsscT0wO3E8czsrK3EpeFtDW3FdXT0KQih0aGlzLDMpO2lmKCF0KXtxPXM7Zm9yKHM9eC5sZW5ndGg7cTxzOysrcSl4W0NbcV1dPTB9UT11KHgpO3Y9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKG0rcCk7cT0wO2ZvcihUPW0rcDtxPFQ7KXN3aXRjaChNPUQodGhpcyxRKSxNKXtjYXNlIDE2OmZvcih6PTMrQih0aGlzLDIpO3otLTspdltxKytdPUY7YnJlYWs7Y2FzZSAxNzpmb3Ioej0zK0IodGhpcywzKTt6LS07KXZbcSsrXT0wO0Y9MDticmVhaztjYXNlIDE4OmZvcih6PTExK0IodGhpcyw3KTt6LS07KXZbcSsrXT0wO0Y9MDticmVhaztkZWZhdWx0OkY9dltxKytdPU19Uj10P3Uodi5zdWJhcnJheSgwLG0pKTp1KHYuc2xpY2UoMCxtKSk7Uz10P3Uodi5zdWJhcnJheShtKSk6dSh2LnNsaWNlKG0pKTt0aGlzLmooUixTKTticmVhaztkZWZhdWx0OnRocm93IEVycm9yKCJ1bmtub3duIEJUWVBFOiAiK2MpO319cmV0dXJuIHRoaXMubSgpfTsKdmFyIEU9WzE2LDE3LDE4LDAsOCw3LDksNiwxMCw1LDExLDQsMTIsMywxMywyLDE0LDEsMTVdLEM9dD9uZXcgVWludDE2QXJyYXkoRSk6RSxHPVszLDQsNSw2LDcsOCw5LDEwLDExLDEzLDE1LDE3LDE5LDIzLDI3LDMxLDM1LDQzLDUxLDU5LDY3LDgzLDk5LDExNSwxMzEsMTYzLDE5NSwyMjcsMjU4LDI1OCwyNThdLEg9dD9uZXcgVWludDE2QXJyYXkoRyk6RyxJPVswLDAsMCwwLDAsMCwwLDAsMSwxLDEsMSwyLDIsMiwyLDMsMywzLDMsNCw0LDQsNCw1LDUsNSw1LDAsMCwwXSxKPXQ/bmV3IFVpbnQ4QXJyYXkoSSk6SSxLPVsxLDIsMyw0LDUsNyw5LDEzLDE3LDI1LDMzLDQ5LDY1LDk3LDEyOSwxOTMsMjU3LDM4NSw1MTMsNzY5LDEwMjUsMTUzNywyMDQ5LDMwNzMsNDA5Nyw2MTQ1LDgxOTMsMTIyODksMTYzODUsMjQ1NzddLEw9dD9uZXcgVWludDE2QXJyYXkoSyk6SyxOPVswLDAsMCwwLDEsMSwyLDIsMywzLDQsNCw1LDUsNiw2LDcsNyw4LDgsOSw5LDEwLDEwLDExLDExLDEyLDEyLDEzLAoxM10sTz10P25ldyBVaW50OEFycmF5KE4pOk4sUD1uZXcgKHQ/VWludDhBcnJheTpBcnJheSkoMjg4KSxVLGRhO1U9MDtmb3IoZGE9UC5sZW5ndGg7VTxkYTsrK1UpUFtVXT0xNDM+PVU/ODoyNTU+PVU/OToyNzk+PVU/Nzo4O3ZhciBiYT11KFApLFY9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKDMwKSxXLGVhO1c9MDtmb3IoZWE9Vi5sZW5ndGg7VzxlYTsrK1cpVltXXT01O3ZhciBjYT11KFYpO2Z1bmN0aW9uIEIoYyxkKXtmb3IodmFyIGE9Yy5mLGI9Yy5jLGU9Yy5pbnB1dCxmPWMuZCxnPWUubGVuZ3RoLGg7YjxkOyl7aWYoZj49Zyl0aHJvdyBFcnJvcigiaW5wdXQgYnVmZmVyIGlzIGJyb2tlbiIpO2F8PWVbZisrXTw8YjtiKz04fWg9YSYoMTw8ZCktMTtjLmY9YT4+PmQ7Yy5jPWItZDtjLmQ9ZjtyZXR1cm4gaH0KZnVuY3Rpb24gRChjLGQpe2Zvcih2YXIgYT1jLmYsYj1jLmMsZT1jLmlucHV0LGY9Yy5kLGc9ZS5sZW5ndGgsaD1kWzBdLGw9ZFsxXSxuLG07YjxsJiYhKGY+PWcpOylhfD1lW2YrK108PGIsYis9ODtuPWhbYSYoMTw8bCktMV07bT1uPj4+MTY7aWYobT5iKXRocm93IEVycm9yKCJpbnZhbGlkIGNvZGUgbGVuZ3RoOiAiK20pO2MuZj1hPj5tO2MuYz1iLW07Yy5kPWY7cmV0dXJuIG4mNjU1MzV9CncucHJvdG90eXBlLmo9ZnVuY3Rpb24oYyxkKXt2YXIgYT10aGlzLmIsYj10aGlzLmE7dGhpcy5uPWM7Zm9yKHZhciBlPWEubGVuZ3RoLTI1OCxmLGcsaCxsOzI1NiE9PShmPUQodGhpcyxjKSk7KWlmKDI1Nj5mKWI+PWUmJih0aGlzLmE9YixhPXRoaXMuZSgpLGI9dGhpcy5hKSxhW2IrK109ZjtlbHNle2c9Zi0yNTc7bD1IW2ddOzA8SltnXSYmKGwrPUIodGhpcyxKW2ddKSk7Zj1EKHRoaXMsZCk7aD1MW2ZdOzA8T1tmXSYmKGgrPUIodGhpcyxPW2ZdKSk7Yj49ZSYmKHRoaXMuYT1iLGE9dGhpcy5lKCksYj10aGlzLmEpO2Zvcig7bC0tOylhW2JdPWFbYisrLWhdfWZvcig7ODw9dGhpcy5jOyl0aGlzLmMtPTgsdGhpcy5kLS07dGhpcy5hPWJ9Owp3LnByb3RvdHlwZS5zPWZ1bmN0aW9uKGMsZCl7dmFyIGE9dGhpcy5iLGI9dGhpcy5hO3RoaXMubj1jO2Zvcih2YXIgZT1hLmxlbmd0aCxmLGcsaCxsOzI1NiE9PShmPUQodGhpcyxjKSk7KWlmKDI1Nj5mKWI+PWUmJihhPXRoaXMuZSgpLGU9YS5sZW5ndGgpLGFbYisrXT1mO2Vsc2V7Zz1mLTI1NztsPUhbZ107MDxKW2ddJiYobCs9Qih0aGlzLEpbZ10pKTtmPUQodGhpcyxkKTtoPUxbZl07MDxPW2ZdJiYoaCs9Qih0aGlzLE9bZl0pKTtiK2w+ZSYmKGE9dGhpcy5lKCksZT1hLmxlbmd0aCk7Zm9yKDtsLS07KWFbYl09YVtiKystaF19Zm9yKDs4PD10aGlzLmM7KXRoaXMuYy09OCx0aGlzLmQtLTt0aGlzLmE9Yn07CncucHJvdG90eXBlLmU9ZnVuY3Rpb24oKXt2YXIgYz1uZXcgKHQ/VWludDhBcnJheTpBcnJheSkodGhpcy5hLTMyNzY4KSxkPXRoaXMuYS0zMjc2OCxhLGIsZT10aGlzLmI7aWYodCljLnNldChlLnN1YmFycmF5KDMyNzY4LGMubGVuZ3RoKSk7ZWxzZXthPTA7Zm9yKGI9Yy5sZW5ndGg7YTxiOysrYSljW2FdPWVbYSszMjc2OF19dGhpcy5nLnB1c2goYyk7dGhpcy5rKz1jLmxlbmd0aDtpZih0KWUuc2V0KGUuc3ViYXJyYXkoZCxkKzMyNzY4KSk7ZWxzZSBmb3IoYT0wOzMyNzY4PmE7KythKWVbYV09ZVtkK2FdO3RoaXMuYT0zMjc2ODtyZXR1cm4gZX07CncucHJvdG90eXBlLnU9ZnVuY3Rpb24oYyl7dmFyIGQsYT10aGlzLmlucHV0Lmxlbmd0aC90aGlzLmQrMXwwLGIsZSxmLGc9dGhpcy5pbnB1dCxoPXRoaXMuYjtjJiYoIm51bWJlciI9PT10eXBlb2YgYy5vJiYoYT1jLm8pLCJudW1iZXIiPT09dHlwZW9mIGMucSYmKGErPWMucSkpOzI+YT8oYj0oZy5sZW5ndGgtdGhpcy5kKS90aGlzLm5bMl0sZj0yNTgqKGIvMil8MCxlPWY8aC5sZW5ndGg/aC5sZW5ndGgrZjpoLmxlbmd0aDw8MSk6ZT1oLmxlbmd0aCphO3Q/KGQ9bmV3IFVpbnQ4QXJyYXkoZSksZC5zZXQoaCkpOmQ9aDtyZXR1cm4gdGhpcy5iPWR9Owp3LnByb3RvdHlwZS5tPWZ1bmN0aW9uKCl7dmFyIGM9MCxkPXRoaXMuYixhPXRoaXMuZyxiLGU9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKHRoaXMuaysodGhpcy5hLTMyNzY4KSksZixnLGgsbDtpZigwPT09YS5sZW5ndGgpcmV0dXJuIHQ/dGhpcy5iLnN1YmFycmF5KDMyNzY4LHRoaXMuYSk6dGhpcy5iLnNsaWNlKDMyNzY4LHRoaXMuYSk7Zj0wO2ZvcihnPWEubGVuZ3RoO2Y8ZzsrK2Ype2I9YVtmXTtoPTA7Zm9yKGw9Yi5sZW5ndGg7aDxsOysraCllW2MrK109YltoXX1mPTMyNzY4O2ZvcihnPXRoaXMuYTtmPGc7KytmKWVbYysrXT1kW2ZdO3RoaXMuZz1bXTtyZXR1cm4gdGhpcy5idWZmZXI9ZX07CncucHJvdG90eXBlLnI9ZnVuY3Rpb24oKXt2YXIgYyxkPXRoaXMuYTt0P3RoaXMucD8oYz1uZXcgVWludDhBcnJheShkKSxjLnNldCh0aGlzLmIuc3ViYXJyYXkoMCxkKSkpOmM9dGhpcy5iLnN1YmFycmF5KDAsZCk6KHRoaXMuYi5sZW5ndGg+ZCYmKHRoaXMuYi5sZW5ndGg9ZCksYz10aGlzLmIpO3JldHVybiB0aGlzLmJ1ZmZlcj1jfTtyKCJabGliLlJhd0luZmxhdGUiLHcpO3IoIlpsaWIuUmF3SW5mbGF0ZS5wcm90b3R5cGUuZGVjb21wcmVzcyIsdy5wcm90b3R5cGUudCk7dmFyIFg9e0FEQVBUSVZFOnksQkxPQ0s6QX0sWSxaLCQsZmE7aWYoT2JqZWN0LmtleXMpWT1PYmplY3Qua2V5cyhYKTtlbHNlIGZvcihaIGluIFk9W10sJD0wLFgpWVskKytdPVo7JD0wO2ZvcihmYT1ZLmxlbmd0aDskPGZhOysrJClaPVlbJF0scigiWmxpYi5SYXdJbmZsYXRlLkJ1ZmZlclR5cGUuIitaLFhbWl0pO30pLmNhbGwodGhpcyk7Cg=="

var bootstrapCode = "Ly8gQ29weXJpZ2h0IDIwMTggVGhlIEdvIEF1dGhvcnMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCi8vIFVzZSBvZiB0aGlzIHNvdXJjZSBjb2RlIGlzIGdvdmVybmVkIGJ5IGEgQlNELXN0eWxlCi8vIGxpY2Vuc2UgdGhhdCBjYW4gYmUgZm91bmQgaW4gdGhlIExJQ0VOU0UgZmlsZS4KCigoKSA9PiB7CgkvLyBNYXAgd2ViIGJyb3dzZXIgQVBJIGFuZCBOb2RlLmpzIEFQSSB0byBhIHNpbmdsZSBjb21tb24gQVBJIChwcmVmZXJyaW5nIHdlYiBzdGFuZGFyZHMgb3ZlciBOb2RlLmpzIEFQSSkuCgljb25zdCBpc05vZGVKUyA9IHR5cGVvZiBwcm9jZXNzICE9PSAidW5kZWZpbmVkIjsKCWlmIChpc05vZGVKUykgewoJCWdsb2JhbC5yZXF1aXJlID0gcmVxdWlyZTsKCQlnbG9iYWwuZnMgPSByZXF1aXJlKCJmcyIpOwoKCQljb25zdCBub2RlQ3J5cHRvID0gcmVxdWlyZSgiY3J5cHRvIik7CgkJZ2xvYmFsLmNyeXB0byA9IHsKCQkJZ2V0UmFuZG9tVmFsdWVzKGIpIHsKCQkJCW5vZGVDcnlwdG8ucmFuZG9tRmlsbFN5bmMoYik7CgkJCX0sCgkJfTsKCgkJZ2xvYmFsLnBlcmZvcm1hbmNlID0gewoJCQlub3coKSB7CgkJCQljb25zdCBbc2VjLCBuc2VjXSA9IHByb2Nlc3MuaHJ0aW1lKCk7CgkJCQlyZXR1cm4gc2VjICogMTAwMCArIG5zZWMgLyAxMDAwMDAwOwoJCQl9LAoJCX07CgoJCWNvbnN0IHV0aWwgPSByZXF1aXJlKCJ1dGlsIik7CgkJZ2xvYmFsLlRleHRFbmNvZGVyID0gdXRpbC5UZXh0RW5jb2RlcjsKCQlnbG9iYWwuVGV4dERlY29kZXIgPSB1dGlsLlRleHREZWNvZGVyOwoJfSBlbHNlIHsKCQlpZiAodHlwZW9mIHdpbmRvdyAhPT0gInVuZGVmaW5lZCIpIHsKCQkJd2luZG93Lmdsb2JhbCA9IHdpbmRvdzsKCQl9IGVsc2UgaWYgKHR5cGVvZiBzZWxmICE9PSAidW5kZWZpbmVkIikgewoJCQlzZWxmLmdsb2JhbCA9IHNlbGY7CgkJfSBlbHNlIHsKCQkJdGhyb3cgbmV3IEVycm9yKCJjYW5ub3QgZXhwb3J0IEdvIChuZWl0aGVyIHdpbmRvdyBub3Igc2VsZiBpcyBkZWZpbmVkKSIpOwoJCX0KCgkJbGV0IG91dHB1dEJ1ZiA9ICIiOwoJCWdsb2JhbC5mcyA9IHsKCQkJY29uc3RhbnRzOiB7IE9fV1JPTkxZOiAtMSwgT19SRFdSOiAtMSwgT19DUkVBVDogLTEsIE9fVFJVTkM6IC0xLCBPX0FQUEVORDogLTEsIE9fRVhDTDogLTEgfSwgLy8gdW51c2VkCgkJCXdyaXRlU3luYyhmZCwgYnVmKSB7CgkJCQlvdXRwdXRCdWYgKz0gZGVjb2Rlci5kZWNvZGUoYnVmKTsKCQkJCWNvbnN0IG5sID0gb3V0cHV0QnVmLmxhc3RJbmRleE9mKCJcbiIpOwoJCQkJaWYgKG5sICE9IC0xKSB7CgkJCQkJY29uc29sZS5sb2cob3V0cHV0QnVmLnN1YnN0cigwLCBubCkpOwoJCQkJCW91dHB1dEJ1ZiA9IG91dHB1dEJ1Zi5zdWJzdHIobmwgKyAxKTsKCQkJCX0KCQkJCXJldHVybiBidWYubGVuZ3RoOwoJCQl9LAoJCQlvcGVuU3luYyhwYXRoLCBmbGFncywgbW9kZSkgewoJCQkJY29uc3QgZXJyID0gbmV3IEVycm9yKCJub3QgaW1wbGVtZW50ZWQiKTsKCQkJCWVyci5jb2RlID0gIkVOT1NZUyI7CgkJCQl0aHJvdyBlcnI7CgkJCX0sCgkJfTsKCX0KCgljb25zdCBlbmNvZGVyID0gbmV3IFRleHRFbmNvZGVyKCJ1dGYtOCIpOwoJY29uc3QgZGVjb2RlciA9IG5ldyBUZXh0RGVjb2RlcigidXRmLTgiKTsKCglnbG9iYWwuR28gPSBjbGFzcyB7CgkJY29uc3RydWN0b3IoKSB7CgkJCXRoaXMuYXJndiA9IFsianMiXTsKCQkJdGhpcy5lbnYgPSB7fTsKCQkJdGhpcy5leGl0ID0gKGNvZGUpID0+IHsKCQkJCWlmIChjb2RlICE9PSAwKSB7CgkJCQkJY29uc29sZS53YXJuKCJleGl0IGNvZGU6IiwgY29kZSk7CgkJCQl9CgkJCX07CgkJCXRoaXMuX2NhbGxiYWNrVGltZW91dHMgPSBuZXcgTWFwKCk7CgkJCXRoaXMuX25leHRDYWxsYmFja1RpbWVvdXRJRCA9IDE7CgoJCQljb25zdCBtZW0gPSAoKSA9PiB7CgkJCQkvLyBUaGUgYnVmZmVyIG1heSBjaGFuZ2Ugd2hlbiByZXF1ZXN0aW5nIG1vcmUgbWVtb3J5LgoJCQkJcmV0dXJuIG5ldyBEYXRhVmlldyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlcik7CgkJCX0KCgkJCWNvbnN0IHNldEludDY0ID0gKGFkZHIsIHYpID0+IHsKCQkJCW1lbSgpLnNldFVpbnQzMihhZGRyICsgMCwgdiwgdHJ1ZSk7CgkJCQltZW0oKS5zZXRVaW50MzIoYWRkciArIDQsIE1hdGguZmxvb3IodiAvIDQyOTQ5NjcyOTYpLCB0cnVlKTsKCQkJfQoKCQkJY29uc3QgZ2V0SW50NjQgPSAoYWRkcikgPT4gewoJCQkJY29uc3QgbG93ID0gbWVtKCkuZ2V0VWludDMyKGFkZHIgKyAwLCB0cnVlKTsKCQkJCWNvbnN0IGhpZ2ggPSBtZW0oKS5nZXRJbnQzMihhZGRyICsgNCwgdHJ1ZSk7CgkJCQlyZXR1cm4gbG93ICsgaGlnaCAqIDQyOTQ5NjcyOTY7CgkJCX0KCgkJCWNvbnN0IGxvYWRWYWx1ZSA9IChhZGRyKSA9PiB7CgkJCQljb25zdCBmID0gbWVtKCkuZ2V0RmxvYXQ2NChhZGRyLCB0cnVlKTsKCQkJCWlmICghaXNOYU4oZikpIHsKCQkJCQlyZXR1cm4gZjsKCQkJCX0KCgkJCQljb25zdCBpZCA9IG1lbSgpLmdldFVpbnQzMihhZGRyLCB0cnVlKTsKCQkJCXJldHVybiB0aGlzLl92YWx1ZXNbaWRdOwoJCQl9CgoJCQljb25zdCBzdG9yZVZhbHVlID0gKGFkZHIsIHYpID0+IHsKCQkJCWNvbnN0IG5hbkhlYWQgPSAweDdGRjgwMDAwOwoKCQkJCWlmICh0eXBlb2YgdiA9PT0gIm51bWJlciIpIHsKCQkJCQlpZiAoaXNOYU4odikpIHsKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIgKyA0LCBuYW5IZWFkLCB0cnVlKTsKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIsIDAsIHRydWUpOwoJCQkJCQlyZXR1cm47CgkJCQkJfQoJCQkJCW1lbSgpLnNldEZsb2F0NjQoYWRkciwgdiwgdHJ1ZSk7CgkJCQkJcmV0dXJuOwoJCQkJfQoKCQkJCXN3aXRjaCAodikgewoJCQkJCWNhc2UgdW5kZWZpbmVkOgoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQsIHRydWUpOwoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciwgMSwgdHJ1ZSk7CgkJCQkJCXJldHVybjsKCQkJCQljYXNlIG51bGw6CgkJCQkJCW1lbSgpLnNldFVpbnQzMihhZGRyICsgNCwgbmFuSGVhZCwgdHJ1ZSk7CgkJCQkJCW1lbSgpLnNldFVpbnQzMihhZGRyLCAyLCB0cnVlKTsKCQkJCQkJcmV0dXJuOwoJCQkJCWNhc2UgdHJ1ZToKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIgKyA0LCBuYW5IZWFkLCB0cnVlKTsKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIsIDMsIHRydWUpOwoJCQkJCQlyZXR1cm47CgkJCQkJY2FzZSBmYWxzZToKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIgKyA0LCBuYW5IZWFkLCB0cnVlKTsKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIsIDQsIHRydWUpOwoJCQkJCQlyZXR1cm47CgkJCQl9CgoJCQkJbGV0IHJlZiA9IHRoaXMuX3JlZnMuZ2V0KHYpOwoJCQkJaWYgKHJlZiA9PT0gdW5kZWZpbmVkKSB7CgkJCQkJcmVmID0gdGhpcy5fdmFsdWVzLmxlbmd0aDsKCQkJCQl0aGlzLl92YWx1ZXMucHVzaCh2KTsKCQkJCQl0aGlzLl9yZWZzLnNldCh2LCByZWYpOwoJCQkJfQoJCQkJbGV0IHR5cGVGbGFnID0gMDsKCQkJCXN3aXRjaCAodHlwZW9mIHYpIHsKCQkJCQljYXNlICJzdHJpbmciOgoJCQkJCQl0eXBlRmxhZyA9IDE7CgkJCQkJCWJyZWFrOwoJCQkJCWNhc2UgInN5bWJvbCI6CgkJCQkJCXR5cGVGbGFnID0gMjsKCQkJCQkJYnJlYWs7CgkJCQkJY2FzZSAiZnVuY3Rpb24iOgoJCQkJCQl0eXBlRmxhZyA9IDM7CgkJCQkJCWJyZWFrOwoJCQkJfQoJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIgKyA0LCBuYW5IZWFkIHwgdHlwZUZsYWcsIHRydWUpOwoJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIsIHJlZiwgdHJ1ZSk7CgkJCX0KCgkJCWNvbnN0IGxvYWRTbGljZSA9IChhZGRyKSA9PiB7CgkJCQljb25zdCBhcnJheSA9IGdldEludDY0KGFkZHIgKyAwKTsKCQkJCWNvbnN0IGxlbiA9IGdldEludDY0KGFkZHIgKyA4KTsKCQkJCXJldHVybiBuZXcgVWludDhBcnJheSh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlciwgYXJyYXksIGxlbik7CgkJCX0KCgkJCWNvbnN0IGxvYWRTbGljZU9mVmFsdWVzID0gKGFkZHIpID0+IHsKCQkJCWNvbnN0IGFycmF5ID0gZ2V0SW50NjQoYWRkciArIDApOwoJCQkJY29uc3QgbGVuID0gZ2V0SW50NjQoYWRkciArIDgpOwoJCQkJY29uc3QgYSA9IG5ldyBBcnJheShsZW4pOwoJCQkJZm9yIChsZXQgaSA9IDA7IGkgPCBsZW47IGkrKykgewoJCQkJCWFbaV0gPSBsb2FkVmFsdWUoYXJyYXkgKyBpICogOCk7CgkJCQl9CgkJCQlyZXR1cm4gYTsKCQkJfQoKCQkJY29uc3QgbG9hZFN0cmluZyA9IChhZGRyKSA9PiB7CgkJCQljb25zdCBzYWRkciA9IGdldEludDY0KGFkZHIgKyAwKTsKCQkJCWNvbnN0IGxlbiA9IGdldEludDY0KGFkZHIgKyA4KTsKCQkJCXJldHVybiBkZWNvZGVyLmRlY29kZShuZXcgRGF0YVZpZXcodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIsIHNhZGRyLCBsZW4pKTsKCQkJfQoKCQkJY29uc3QgdGltZU9yaWdpbiA9IERhdGUubm93KCkgLSBwZXJmb3JtYW5jZS5ub3coKTsKCQkJdGhpcy5pbXBvcnRPYmplY3QgPSB7CgkJCQlnbzogewoJCQkJCS8vIGZ1bmMgd2FzbUV4aXQoY29kZSBpbnQzMikKCQkJCQkicnVudGltZS53YXNtRXhpdCI6IChzcCkgPT4gewoJCQkJCQljb25zdCBjb2RlID0gbWVtKCkuZ2V0SW50MzIoc3AgKyA4LCB0cnVlKTsKCQkJCQkJdGhpcy5leGl0ZWQgPSB0cnVlOwoJCQkJCQlkZWxldGUgdGhpcy5faW5zdDsKCQkJCQkJZGVsZXRlIHRoaXMuX3ZhbHVlczsKCQkJCQkJZGVsZXRlIHRoaXMuX3JlZnM7CgkJCQkJCXRoaXMuZXhpdChjb2RlKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHdhc21Xcml0ZShmZCB1aW50cHRyLCBwIHVuc2FmZS5Qb2ludGVyLCBuIGludDMyKQoJCQkJCSJydW50aW1lLndhc21Xcml0ZSI6IChzcCkgPT4gewoJCQkJCQljb25zdCBmZCA9IGdldEludDY0KHNwICsgOCk7CgkJCQkJCWNvbnN0IHAgPSBnZXRJbnQ2NChzcCArIDE2KTsKCQkJCQkJY29uc3QgbiA9IG1lbSgpLmdldEludDMyKHNwICsgMjQsIHRydWUpOwoJCQkJCQlmcy53cml0ZVN5bmMoZmQsIG5ldyBVaW50OEFycmF5KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyLCBwLCBuKSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyBuYW5vdGltZSgpIGludDY0CgkJCQkJInJ1bnRpbWUubmFub3RpbWUiOiAoc3ApID0+IHsKCQkJCQkJc2V0SW50NjQoc3AgKyA4LCAodGltZU9yaWdpbiArIHBlcmZvcm1hbmNlLm5vdygpKSAqIDEwMDAwMDApOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgd2FsbHRpbWUoKSAoc2VjIGludDY0LCBuc2VjIGludDMyKQoJCQkJCSJydW50aW1lLndhbGx0aW1lIjogKHNwKSA9PiB7CgkJCQkJCWNvbnN0IG1zZWMgPSAobmV3IERhdGUpLmdldFRpbWUoKTsKCQkJCQkJc2V0SW50NjQoc3AgKyA4LCBtc2VjIC8gMTAwMCk7CgkJCQkJCW1lbSgpLnNldEludDMyKHNwICsgMTYsIChtc2VjICUgMTAwMCkgKiAxMDAwMDAwLCB0cnVlKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHNjaGVkdWxlQ2FsbGJhY2soZGVsYXkgaW50NjQpIGludDMyCgkJCQkJInJ1bnRpbWUuc2NoZWR1bGVDYWxsYmFjayI6IChzcCkgPT4gewoJCQkJCQljb25zdCBpZCA9IHRoaXMuX25leHRDYWxsYmFja1RpbWVvdXRJRDsKCQkJCQkJdGhpcy5fbmV4dENhbGxiYWNrVGltZW91dElEKys7CgkJCQkJCXRoaXMuX2NhbGxiYWNrVGltZW91dHMuc2V0KGlkLCBzZXRUaW1lb3V0KAoJCQkJCQkJKCkgPT4geyB0aGlzLl9yZXNvbHZlQ2FsbGJhY2tQcm9taXNlKCk7IH0sCgkJCQkJCQlnZXRJbnQ2NChzcCArIDgpICsgMSwgLy8gc2V0VGltZW91dCBoYXMgYmVlbiBzZWVuIHRvIGZpcmUgdXAgdG8gMSBtaWxsaXNlY29uZCBlYXJseQoJCQkJCQkpKTsKCQkJCQkJbWVtKCkuc2V0SW50MzIoc3AgKyAxNiwgaWQsIHRydWUpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgY2xlYXJTY2hlZHVsZWRDYWxsYmFjayhpZCBpbnQzMikKCQkJCQkicnVudGltZS5jbGVhclNjaGVkdWxlZENhbGxiYWNrIjogKHNwKSA9PiB7CgkJCQkJCWNvbnN0IGlkID0gbWVtKCkuZ2V0SW50MzIoc3AgKyA4LCB0cnVlKTsKCQkJCQkJY2xlYXJUaW1lb3V0KHRoaXMuX2NhbGxiYWNrVGltZW91dHMuZ2V0KGlkKSk7CgkJCQkJCXRoaXMuX2NhbGxiYWNrVGltZW91dHMuZGVsZXRlKGlkKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIGdldFJhbmRvbURhdGEociBbXWJ5dGUpCgkJCQkJInJ1bnRpbWUuZ2V0UmFuZG9tRGF0YSI6IChzcCkgPT4gewoJCQkJCQljcnlwdG8uZ2V0UmFuZG9tVmFsdWVzKGxvYWRTbGljZShzcCArIDgpKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHN0cmluZ1ZhbCh2YWx1ZSBzdHJpbmcpIHJlZgoJCQkJCSJzeXNjYWxsL2pzLnN0cmluZ1ZhbCI6IChzcCkgPT4gewoJCQkJCQlzdG9yZVZhbHVlKHNwICsgMjQsIGxvYWRTdHJpbmcoc3AgKyA4KSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB2YWx1ZUdldCh2IHJlZiwgcCBzdHJpbmcpIHJlZgoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlR2V0IjogKHNwKSA9PiB7CgkJCQkJCXN0b3JlVmFsdWUoc3AgKyAzMiwgUmVmbGVjdC5nZXQobG9hZFZhbHVlKHNwICsgOCksIGxvYWRTdHJpbmcoc3AgKyAxNikpKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHZhbHVlU2V0KHYgcmVmLCBwIHN0cmluZywgeCByZWYpCgkJCQkJInN5c2NhbGwvanMudmFsdWVTZXQiOiAoc3ApID0+IHsKCQkJCQkJUmVmbGVjdC5zZXQobG9hZFZhbHVlKHNwICsgOCksIGxvYWRTdHJpbmcoc3AgKyAxNiksIGxvYWRWYWx1ZShzcCArIDMyKSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB2YWx1ZUluZGV4KHYgcmVmLCBpIGludCkgcmVmCgkJCQkJInN5c2NhbGwvanMudmFsdWVJbmRleCI6IChzcCkgPT4gewoJCQkJCQlzdG9yZVZhbHVlKHNwICsgMjQsIFJlZmxlY3QuZ2V0KGxvYWRWYWx1ZShzcCArIDgpLCBnZXRJbnQ2NChzcCArIDE2KSkpOwoJCQkJCX0sCgoJCQkJCS8vIHZhbHVlU2V0SW5kZXgodiByZWYsIGkgaW50LCB4IHJlZikKCQkJCQkic3lzY2FsbC9qcy52YWx1ZVNldEluZGV4IjogKHNwKSA9PiB7CgkJCQkJCVJlZmxlY3Quc2V0KGxvYWRWYWx1ZShzcCArIDgpLCBnZXRJbnQ2NChzcCArIDE2KSwgbG9hZFZhbHVlKHNwICsgMjQpKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHZhbHVlQ2FsbCh2IHJlZiwgbSBzdHJpbmcsIGFyZ3MgW11yZWYpIChyZWYsIGJvb2wpCgkJCQkJInN5c2NhbGwvanMudmFsdWVDYWxsIjogKHNwKSA9PiB7CgkJCQkJCXRyeSB7CgkJCQkJCQljb25zdCB2ID0gbG9hZFZhbHVlKHNwICsgOCk7CgkJCQkJCQljb25zdCBtID0gUmVmbGVjdC5nZXQodiwgbG9hZFN0cmluZyhzcCArIDE2KSk7CgkJCQkJCQljb25zdCBhcmdzID0gbG9hZFNsaWNlT2ZWYWx1ZXMoc3AgKyAzMik7CgkJCQkJCQlzdG9yZVZhbHVlKHNwICsgNTYsIFJlZmxlY3QuYXBwbHkobSwgdiwgYXJncykpOwoJCQkJCQkJbWVtKCkuc2V0VWludDgoc3AgKyA2NCwgMSk7CgkJCQkJCX0gY2F0Y2ggKGVycikgewoJCQkJCQkJc3RvcmVWYWx1ZShzcCArIDU2LCBlcnIpOwoJCQkJCQkJbWVtKCkuc2V0VWludDgoc3AgKyA2NCwgMCk7CgkJCQkJCX0KCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHZhbHVlSW52b2tlKHYgcmVmLCBhcmdzIFtdcmVmKSAocmVmLCBib29sKQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlSW52b2tlIjogKHNwKSA9PiB7CgkJCQkJCXRyeSB7CgkJCQkJCQljb25zdCB2ID0gbG9hZFZhbHVlKHNwICsgOCk7CgkJCQkJCQljb25zdCBhcmdzID0gbG9hZFNsaWNlT2ZWYWx1ZXMoc3AgKyAxNik7CgkJCQkJCQlzdG9yZVZhbHVlKHNwICsgNDAsIFJlZmxlY3QuYXBwbHkodiwgdW5kZWZpbmVkLCBhcmdzKSk7CgkJCQkJCQltZW0oKS5zZXRVaW50OChzcCArIDQ4LCAxKTsKCQkJCQkJfSBjYXRjaCAoZXJyKSB7CgkJCQkJCQlzdG9yZVZhbHVlKHNwICsgNDAsIGVycik7CgkJCQkJCQltZW0oKS5zZXRVaW50OChzcCArIDQ4LCAwKTsKCQkJCQkJfQoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVOZXcodiByZWYsIGFyZ3MgW11yZWYpIChyZWYsIGJvb2wpCgkJCQkJInN5c2NhbGwvanMudmFsdWVOZXciOiAoc3ApID0+IHsKCQkJCQkJdHJ5IHsKCQkJCQkJCWNvbnN0IHYgPSBsb2FkVmFsdWUoc3AgKyA4KTsKCQkJCQkJCWNvbnN0IGFyZ3MgPSBsb2FkU2xpY2VPZlZhbHVlcyhzcCArIDE2KTsKCQkJCQkJCXN0b3JlVmFsdWUoc3AgKyA0MCwgUmVmbGVjdC5jb25zdHJ1Y3QodiwgYXJncykpOwoJCQkJCQkJbWVtKCkuc2V0VWludDgoc3AgKyA0OCwgMSk7CgkJCQkJCX0gY2F0Y2ggKGVycikgewoJCQkJCQkJc3RvcmVWYWx1ZShzcCArIDQwLCBlcnIpOwoJCQkJCQkJbWVtKCkuc2V0VWludDgoc3AgKyA0OCwgMCk7CgkJCQkJCX0KCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHZhbHVlTGVuZ3RoKHYgcmVmKSBpbnQKCQkJCQkic3lzY2FsbC9qcy52YWx1ZUxlbmd0aCI6IChzcCkgPT4gewoJCQkJCQlzZXRJbnQ2NChzcCArIDE2LCBwYXJzZUludChsb2FkVmFsdWUoc3AgKyA4KS5sZW5ndGgpKTsKCQkJCQl9LAoKCQkJCQkvLyB2YWx1ZVByZXBhcmVTdHJpbmcodiByZWYpIChyZWYsIGludCkKCQkJCQkic3lzY2FsbC9qcy52YWx1ZVByZXBhcmVTdHJpbmciOiAoc3ApID0+IHsKCQkJCQkJY29uc3Qgc3RyID0gZW5jb2Rlci5lbmNvZGUoU3RyaW5nKGxvYWRWYWx1ZShzcCArIDgpKSk7CgkJCQkJCXN0b3JlVmFsdWUoc3AgKyAxNiwgc3RyKTsKCQkJCQkJc2V0SW50NjQoc3AgKyAyNCwgc3RyLmxlbmd0aCk7CgkJCQkJfSwKCgkJCQkJLy8gdmFsdWVMb2FkU3RyaW5nKHYgcmVmLCBiIFtdYnl0ZSkKCQkJCQkic3lzY2FsbC9qcy52YWx1ZUxvYWRTdHJpbmciOiAoc3ApID0+IHsKCQkJCQkJY29uc3Qgc3RyID0gbG9hZFZhbHVlKHNwICsgOCk7CgkJCQkJCWxvYWRTbGljZShzcCArIDE2KS5zZXQoc3RyKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHZhbHVlSW5zdGFuY2VPZih2IHJlZiwgdCByZWYpIGJvb2wKCQkJCQkic3lzY2FsbC9qcy52YWx1ZUluc3RhbmNlT2YiOiAoc3ApID0+IHsKCQkJCQkJbWVtKCkuc2V0VWludDgoc3AgKyAyNCwgbG9hZFZhbHVlKHNwICsgOCkgaW5zdGFuY2VvZiBsb2FkVmFsdWUoc3AgKyAxNikpOwoJCQkJCX0sCgoJCQkJCSJkZWJ1ZyI6ICh2YWx1ZSkgPT4gewoJCQkJCQljb25zb2xlLmxvZyh2YWx1ZSk7CgkJCQkJfSwKCQkJCX0KCQkJfTsKCQl9CgoJCWFzeW5jIHJ1bihpbnN0YW5jZSkgewoJCQl0aGlzLl9pbnN0ID0gaW5zdGFuY2U7CgkJCXRoaXMuX3ZhbHVlcyA9IFsgLy8gVE9ETzogZ2FyYmFnZSBjb2xsZWN0aW9uCgkJCQlOYU4sCgkJCQl1bmRlZmluZWQsCgkJCQludWxsLAoJCQkJdHJ1ZSwKCQkJCWZhbHNlLAoJCQkJZ2xvYmFsLAoJCQkJdGhpcy5faW5zdC5leHBvcnRzLm1lbSwKCQkJCXRoaXMsCgkJCV07CgkJCXRoaXMuX3JlZnMgPSBuZXcgTWFwKCk7CgkJCXRoaXMuX2NhbGxiYWNrU2h1dGRvd24gPSBmYWxzZTsKCQkJdGhpcy5leGl0ZWQgPSBmYWxzZTsKCgkJCWNvbnN0IG1lbSA9IG5ldyBEYXRhVmlldyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlcikKCgkJCS8vIFBhc3MgY29tbWFuZCBsaW5lIGFyZ3VtZW50cyBhbmQgZW52aXJvbm1lbnQgdmFyaWFibGVzIHRvIFdlYkFzc2VtYmx5IGJ5IHdyaXRpbmcgdGhlbSB0byB0aGUgbGluZWFyIG1lbW9yeS4KCQkJbGV0IG9mZnNldCA9IDQwOTY7CgoJCQljb25zdCBzdHJQdHIgPSAoc3RyKSA9PiB7CgkJCQlsZXQgcHRyID0gb2Zmc2V0OwoJCQkJbmV3IFVpbnQ4QXJyYXkobWVtLmJ1ZmZlciwgb2Zmc2V0LCBzdHIubGVuZ3RoICsgMSkuc2V0KGVuY29kZXIuZW5jb2RlKHN0ciArICJcMCIpKTsKCQkJCW9mZnNldCArPSBzdHIubGVuZ3RoICsgKDggLSAoc3RyLmxlbmd0aCAlIDgpKTsKCQkJCXJldHVybiBwdHI7CgkJCX07CgoJCQljb25zdCBhcmdjID0gdGhpcy5hcmd2Lmxlbmd0aDsKCgkJCWNvbnN0IGFyZ3ZQdHJzID0gW107CgkJCXRoaXMuYXJndi5mb3JFYWNoKChhcmcpID0+IHsKCQkJCWFyZ3ZQdHJzLnB1c2goc3RyUHRyKGFyZykpOwoJCQl9KTsKCgkJCWNvbnN0IGtleXMgPSBPYmplY3Qua2V5cyh0aGlzLmVudikuc29ydCgpOwoJCQlhcmd2UHRycy5wdXNoKGtleXMubGVuZ3RoKTsKCQkJa2V5cy5mb3JFYWNoKChrZXkpID0+IHsKCQkJCWFyZ3ZQdHJzLnB1c2goc3RyUHRyKGAke2tleX09JHt0aGlzLmVudltrZXldfWApKTsKCQkJfSk7CgoJCQljb25zdCBhcmd2ID0gb2Zmc2V0OwoJCQlhcmd2UHRycy5mb3JFYWNoKChwdHIpID0+IHsKCQkJCW1lbS5zZXRVaW50MzIob2Zmc2V0LCBwdHIsIHRydWUpOwoJCQkJbWVtLnNldFVpbnQzMihvZmZzZXQgKyA0LCAwLCB0cnVlKTsKCQkJCW9mZnNldCArPSA4OwoJCQl9KTsKCgkJCXdoaWxlICh0cnVlKSB7CgkJCQljb25zdCBjYWxsYmFja1Byb21pc2UgPSBuZXcgUHJvbWlzZSgocmVzb2x2ZSkgPT4gewoJCQkJCXRoaXMuX3Jlc29sdmVDYWxsYmFja1Byb21pc2UgPSAoKSA9PiB7CgkJCQkJCWlmICh0aGlzLmV4aXRlZCkgewoJCQkJCQkJdGhyb3cgbmV3IEVycm9yKCJiYWQgY2FsbGJhY2s6IEdvIHByb2dyYW0gaGFzIGFscmVhZHkgZXhpdGVkIik7CgkJCQkJCX0KCQkJCQkJc2V0VGltZW91dChyZXNvbHZlLCAwKTsgLy8gbWFrZSBzdXJlIGl0IGlzIGFzeW5jaHJvbm91cwoJCQkJCX07CgkJCQl9KTsKCQkJCXRoaXMuX2luc3QuZXhwb3J0cy5ydW4oYXJnYywgYXJndik7CgkJCQlpZiAodGhpcy5leGl0ZWQpIHsKCQkJCQlicmVhazsKCQkJCX0KCQkJCWF3YWl0IGNhbGxiYWNrUHJvbWlzZTsKCQkJfQoJCX0KCgkJc3RhdGljIF9tYWtlQ2FsbGJhY2tIZWxwZXIoaWQsIHBlbmRpbmdDYWxsYmFja3MsIGdvKSB7CgkJCXJldHVybiBmdW5jdGlvbigpIHsKCQkJCXBlbmRpbmdDYWxsYmFja3MucHVzaCh7IGlkOiBpZCwgYXJnczogYXJndW1lbnRzIH0pOwoJCQkJZ28uX3Jlc29sdmVDYWxsYmFja1Byb21pc2UoKTsKCQkJfTsKCQl9CgoJCXN0YXRpYyBfbWFrZUV2ZW50Q2FsbGJhY2tIZWxwZXIocHJldmVudERlZmF1bHQsIHN0b3BQcm9wYWdhdGlvbiwgc3RvcEltbWVkaWF0ZVByb3BhZ2F0aW9uLCBmbikgewoJCQlyZXR1cm4gZnVuY3Rpb24oZXZlbnQpIHsKCQkJCWlmIChwcmV2ZW50RGVmYXVsdCkgewoJCQkJCWV2ZW50LnByZXZlbnREZWZhdWx0KCk7CgkJCQl9CgkJCQlpZiAoc3RvcFByb3BhZ2F0aW9uKSB7CgkJCQkJZXZlbnQuc3RvcFByb3BhZ2F0aW9uKCk7CgkJCQl9CgkJCQlpZiAoc3RvcEltbWVkaWF0ZVByb3BhZ2F0aW9uKSB7CgkJCQkJZXZlbnQuc3RvcEltbWVkaWF0ZVByb3BhZ2F0aW9uKCk7CgkJCQl9CgkJCQlmbihldmVudCk7CgkJCX07CgkJfQoJfQoKCWlmIChpc05vZGVKUykgewoJCWlmIChwcm9jZXNzLmFyZ3YubGVuZ3RoIDwgMykgewoJCQlwcm9jZXNzLnN0ZGVyci53cml0ZSgidXNhZ2U6IGdvX2pzX3dhc21fZXhlYyBbd2FzbSBiaW5hcnldIFthcmd1bWVudHNdXG4iKTsKCQkJcHJvY2Vzcy5leGl0KDEpOwoJCX0KCgkJY29uc3QgZ28gPSBuZXcgR28oKTsKCQlnby5hcmd2ID0gcHJvY2Vzcy5hcmd2LnNsaWNlKDIpOwoJCWdvLmVudiA9IHByb2Nlc3MuZW52OwoJCWdvLmV4aXQgPSBwcm9jZXNzLmV4aXQ7CgkJV2ViQXNzZW1ibHkuaW5zdGFudGlhdGUoZnMucmVhZEZpbGVTeW5jKHByb2Nlc3MuYXJndlsyXSksIGdvLmltcG9ydE9iamVjdCkudGhlbigocmVzdWx0KSA9PiB7CgkJCXByb2Nlc3Mub24oImV4aXQiLCAoY29kZSkgPT4geyAvLyBOb2RlLmpzIGV4aXRzIGlmIG5vIGNhbGxiYWNrIGlzIHBlbmRpbmcKCQkJCWlmIChjb2RlID09PSAwICYmICFnby5leGl0ZWQpIHsKCQkJCQkvLyBkZWFkbG9jaywgbWFrZSBHbyBwcmludCBlcnJvciBhbmQgc3RhY2sgdHJhY2VzCgkJCQkJZ28uX2NhbGxiYWNrU2h1dGRvd24gPSB0cnVlOwoJCQkJCWdvLl9pbnN0LmV4cG9ydHMucnVuKCk7CgkJCQl9CgkJCX0pOwoJCQlyZXR1cm4gZ28ucnVuKHJlc3VsdC5pbnN0YW5jZSk7CgkJfSkuY2F0Y2goKGVycikgPT4gewoJCQl0aHJvdyBlcnI7CgkJfSk7Cgl9Cn0pKCk7Cg=="