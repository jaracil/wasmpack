package main

var base64Code = "KGZ1bmN0aW9uKHIpe2lmKHR5cGVvZiBleHBvcnRzPT09Im9iamVjdCImJnR5cGVvZiBtb2R1bGUhPT0idW5kZWZpbmVkIil7bW9kdWxlLmV4cG9ydHM9cigpfWVsc2UgaWYodHlwZW9mIGRlZmluZT09PSJmdW5jdGlvbiImJmRlZmluZS5hbWQpe2RlZmluZShbXSxyKX1lbHNle3ZhciBlO2lmKHR5cGVvZiB3aW5kb3chPT0idW5kZWZpbmVkIil7ZT13aW5kb3d9ZWxzZSBpZih0eXBlb2YgZ2xvYmFsIT09InVuZGVmaW5lZCIpe2U9Z2xvYmFsfWVsc2UgaWYodHlwZW9mIHNlbGYhPT0idW5kZWZpbmVkIil7ZT1zZWxmfWVsc2V7ZT10aGlzfWUuYmFzZTY0anM9cigpfX0pKGZ1bmN0aW9uKCl7dmFyIHIsZSxuO3JldHVybiBmdW5jdGlvbigpe2Z1bmN0aW9uIHIoZSxuLHQpe2Z1bmN0aW9uIG8oZixpKXtpZighbltmXSl7aWYoIWVbZl0pe3ZhciB1PSJmdW5jdGlvbiI9PXR5cGVvZiByZXF1aXJlJiZyZXF1aXJlO2lmKCFpJiZ1KXJldHVybiB1KGYsITApO2lmKGEpcmV0dXJuIGEoZiwhMCk7dmFyIHY9bmV3IEVycm9yKCJDYW5ub3QgZmluZCBtb2R1bGUgJyIrZisiJyIpO3Rocm93IHYuY29kZT0iTU9EVUxFX05PVF9GT1VORCIsdn12YXIgZD1uW2ZdPXtleHBvcnRzOnt9fTtlW2ZdWzBdLmNhbGwoZC5leHBvcnRzLGZ1bmN0aW9uKHIpe3ZhciBuPWVbZl1bMV1bcl07cmV0dXJuIG8obnx8cil9LGQsZC5leHBvcnRzLHIsZSxuLHQpfXJldHVybiBuW2ZdLmV4cG9ydHN9Zm9yKHZhciBhPSJmdW5jdGlvbiI9PXR5cGVvZiByZXF1aXJlJiZyZXF1aXJlLGY9MDtmPHQubGVuZ3RoO2YrKylvKHRbZl0pO3JldHVybiBvfXJldHVybiByfSgpKHsiLyI6W2Z1bmN0aW9uKHIsZSxuKXsidXNlIHN0cmljdCI7bi5ieXRlTGVuZ3RoPWQ7bi50b0J5dGVBcnJheT1oO24uZnJvbUJ5dGVBcnJheT1wO3ZhciB0PVtdO3ZhciBvPVtdO3ZhciBhPXR5cGVvZiBVaW50OEFycmF5IT09InVuZGVmaW5lZCI/VWludDhBcnJheTpBcnJheTt2YXIgZj0iQUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODkrLyI7Zm9yKHZhciBpPTAsdT1mLmxlbmd0aDtpPHU7KytpKXt0W2ldPWZbaV07b1tmLmNoYXJDb2RlQXQoaSldPWl9b1siLSIuY2hhckNvZGVBdCgwKV09NjI7b1siXyIuY2hhckNvZGVBdCgwKV09NjM7ZnVuY3Rpb24gdihyKXt2YXIgZT1yLmxlbmd0aDtpZihlJTQ+MCl7dGhyb3cgbmV3IEVycm9yKCJJbnZhbGlkIHN0cmluZy4gTGVuZ3RoIG11c3QgYmUgYSBtdWx0aXBsZSBvZiA0Iil9dmFyIG49ci5pbmRleE9mKCI9Iik7aWYobj09PS0xKW49ZTt2YXIgdD1uPT09ZT8wOjQtbiU0O3JldHVybltuLHRdfWZ1bmN0aW9uIGQocil7dmFyIGU9dihyKTt2YXIgbj1lWzBdO3ZhciB0PWVbMV07cmV0dXJuKG4rdCkqMy80LXR9ZnVuY3Rpb24gYyhyLGUsbil7cmV0dXJuKGUrbikqMy80LW59ZnVuY3Rpb24gaChyKXt2YXIgZTt2YXIgbj12KHIpO3ZhciB0PW5bMF07dmFyIGY9blsxXTt2YXIgaT1uZXcgYShjKHIsdCxmKSk7dmFyIHU9MDt2YXIgZD1mPjA/dC00OnQ7Zm9yKHZhciBoPTA7aDxkO2grPTQpe2U9b1tyLmNoYXJDb2RlQXQoaCldPDwxOHxvW3IuY2hhckNvZGVBdChoKzEpXTw8MTJ8b1tyLmNoYXJDb2RlQXQoaCsyKV08PDZ8b1tyLmNoYXJDb2RlQXQoaCszKV07aVt1KytdPWU+PjE2JjI1NTtpW3UrK109ZT4+OCYyNTU7aVt1KytdPWUmMjU1fWlmKGY9PT0yKXtlPW9bci5jaGFyQ29kZUF0KGgpXTw8MnxvW3IuY2hhckNvZGVBdChoKzEpXT4+NDtpW3UrK109ZSYyNTV9aWYoZj09PTEpe2U9b1tyLmNoYXJDb2RlQXQoaCldPDwxMHxvW3IuY2hhckNvZGVBdChoKzEpXTw8NHxvW3IuY2hhckNvZGVBdChoKzIpXT4+MjtpW3UrK109ZT4+OCYyNTU7aVt1KytdPWUmMjU1fXJldHVybiBpfWZ1bmN0aW9uIHMocil7cmV0dXJuIHRbcj4+MTgmNjNdK3Rbcj4+MTImNjNdK3Rbcj4+NiY2M10rdFtyJjYzXX1mdW5jdGlvbiBsKHIsZSxuKXt2YXIgdDt2YXIgbz1bXTtmb3IodmFyIGE9ZTthPG47YSs9Myl7dD0oclthXTw8MTYmMTY3MTE2ODApKyhyW2ErMV08PDgmNjUyODApKyhyW2ErMl0mMjU1KTtvLnB1c2gocyh0KSl9cmV0dXJuIG8uam9pbigiIil9ZnVuY3Rpb24gcChyKXt2YXIgZTt2YXIgbj1yLmxlbmd0aDt2YXIgbz1uJTM7dmFyIGE9W107dmFyIGY9MTYzODM7Zm9yKHZhciBpPTAsdT1uLW87aTx1O2krPWYpe2EucHVzaChsKHIsaSxpK2Y+dT91OmkrZikpfWlmKG89PT0xKXtlPXJbbi0xXTthLnB1c2godFtlPj4yXSt0W2U8PDQmNjNdKyI9PSIpfWVsc2UgaWYobz09PTIpe2U9KHJbbi0yXTw8OCkrcltuLTFdO2EucHVzaCh0W2U+PjEwXSt0W2U+PjQmNjNdK3RbZTw8MiY2M10rIj0iKX1yZXR1cm4gYS5qb2luKCIiKX19LHt9XX0se30sW10pKCIvIil9KTsK"

var inflateCode = "LyoqIEBsaWNlbnNlIHpsaWIuanMgMjAxMiAtIGltYXlhIFsgaHR0cHM6Ly9naXRodWIuY29tL2ltYXlhL3psaWIuanMgXSBUaGUgTUlUIExpY2Vuc2UgKi8oZnVuY3Rpb24oKSB7J3VzZSBzdHJpY3QnO3ZhciBrPXZvaWQgMCxhYT10aGlzO2Z1bmN0aW9uIHIoYyxkKXt2YXIgYT1jLnNwbGl0KCIuIiksYj1hYTshKGFbMF1pbiBiKSYmYi5leGVjU2NyaXB0JiZiLmV4ZWNTY3JpcHQoInZhciAiK2FbMF0pO2Zvcih2YXIgZTthLmxlbmd0aCYmKGU9YS5zaGlmdCgpKTspIWEubGVuZ3RoJiZkIT09az9iW2VdPWQ6Yj1iW2VdP2JbZV06YltlXT17fX07dmFyIHQ9InVuZGVmaW5lZCIhPT10eXBlb2YgVWludDhBcnJheSYmInVuZGVmaW5lZCIhPT10eXBlb2YgVWludDE2QXJyYXkmJiJ1bmRlZmluZWQiIT09dHlwZW9mIFVpbnQzMkFycmF5JiYidW5kZWZpbmVkIiE9PXR5cGVvZiBEYXRhVmlldztmdW5jdGlvbiB1KGMpe3ZhciBkPWMubGVuZ3RoLGE9MCxiPU51bWJlci5QT1NJVElWRV9JTkZJTklUWSxlLGYsZyxoLGwsbixtLHAscyx4O2ZvcihwPTA7cDxkOysrcCljW3BdPmEmJihhPWNbcF0pLGNbcF08YiYmKGI9Y1twXSk7ZT0xPDxhO2Y9bmV3ICh0P1VpbnQzMkFycmF5OkFycmF5KShlKTtnPTE7aD0wO2ZvcihsPTI7Zzw9YTspe2ZvcihwPTA7cDxkOysrcClpZihjW3BdPT09Zyl7bj0wO209aDtmb3Iocz0wO3M8ZzsrK3Mpbj1uPDwxfG0mMSxtPj49MTt4PWc8PDE2fHA7Zm9yKHM9bjtzPGU7cys9bClmW3NdPXg7KytofSsrZztoPDw9MTtsPDw9MX1yZXR1cm5bZixhLGJdfTtmdW5jdGlvbiB3KGMsZCl7dGhpcy5nPVtdO3RoaXMuaD0zMjc2ODt0aGlzLmM9dGhpcy5mPXRoaXMuZD10aGlzLms9MDt0aGlzLmlucHV0PXQ/bmV3IFVpbnQ4QXJyYXkoYyk6Yzt0aGlzLmw9ITE7dGhpcy5pPXk7dGhpcy5wPSExO2lmKGR8fCEoZD17fSkpZC5pbmRleCYmKHRoaXMuZD1kLmluZGV4KSxkLmJ1ZmZlclNpemUmJih0aGlzLmg9ZC5idWZmZXJTaXplKSxkLmJ1ZmZlclR5cGUmJih0aGlzLmk9ZC5idWZmZXJUeXBlKSxkLnJlc2l6ZSYmKHRoaXMucD1kLnJlc2l6ZSk7c3dpdGNoKHRoaXMuaSl7Y2FzZSBBOnRoaXMuYT0zMjc2ODt0aGlzLmI9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKDMyNzY4K3RoaXMuaCsyNTgpO2JyZWFrO2Nhc2UgeTp0aGlzLmE9MDt0aGlzLmI9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKHRoaXMuaCk7dGhpcy5lPXRoaXMudTt0aGlzLm09dGhpcy5yO3RoaXMuaj10aGlzLnM7YnJlYWs7ZGVmYXVsdDp0aHJvdyBFcnJvcigiaW52YWxpZCBpbmZsYXRlIG1vZGUiKTsKfX12YXIgQT0wLHk9MTsKdy5wcm90b3R5cGUudD1mdW5jdGlvbigpe2Zvcig7IXRoaXMubDspe3ZhciBjPUIodGhpcywzKTtjJjEmJih0aGlzLmw9ITApO2M+Pj49MTtzd2l0Y2goYyl7Y2FzZSAwOnZhciBkPXRoaXMuaW5wdXQsYT10aGlzLmQsYj10aGlzLmIsZT10aGlzLmEsZj1kLmxlbmd0aCxnPWssaD1rLGw9Yi5sZW5ndGgsbj1rO3RoaXMuYz10aGlzLmY9MDtpZihhKzE+PWYpdGhyb3cgRXJyb3IoImludmFsaWQgdW5jb21wcmVzc2VkIGJsb2NrIGhlYWRlcjogTEVOIik7Zz1kW2ErK118ZFthKytdPDw4O2lmKGErMT49Zil0aHJvdyBFcnJvcigiaW52YWxpZCB1bmNvbXByZXNzZWQgYmxvY2sgaGVhZGVyOiBOTEVOIik7aD1kW2ErK118ZFthKytdPDw4O2lmKGc9PT1+aCl0aHJvdyBFcnJvcigiaW52YWxpZCB1bmNvbXByZXNzZWQgYmxvY2sgaGVhZGVyOiBsZW5ndGggdmVyaWZ5Iik7aWYoYStnPmQubGVuZ3RoKXRocm93IEVycm9yKCJpbnB1dCBidWZmZXIgaXMgYnJva2VuIik7c3dpdGNoKHRoaXMuaSl7Y2FzZSBBOmZvcig7ZStnPgpiLmxlbmd0aDspe249bC1lO2ctPW47aWYodCliLnNldChkLnN1YmFycmF5KGEsYStuKSxlKSxlKz1uLGErPW47ZWxzZSBmb3IoO24tLTspYltlKytdPWRbYSsrXTt0aGlzLmE9ZTtiPXRoaXMuZSgpO2U9dGhpcy5hfWJyZWFrO2Nhc2UgeTpmb3IoO2UrZz5iLmxlbmd0aDspYj10aGlzLmUoe286Mn0pO2JyZWFrO2RlZmF1bHQ6dGhyb3cgRXJyb3IoImludmFsaWQgaW5mbGF0ZSBtb2RlIik7fWlmKHQpYi5zZXQoZC5zdWJhcnJheShhLGErZyksZSksZSs9ZyxhKz1nO2Vsc2UgZm9yKDtnLS07KWJbZSsrXT1kW2ErK107dGhpcy5kPWE7dGhpcy5hPWU7dGhpcy5iPWI7YnJlYWs7Y2FzZSAxOnRoaXMuaihiYSxjYSk7YnJlYWs7Y2FzZSAyOmZvcih2YXIgbT1CKHRoaXMsNSkrMjU3LHA9Qih0aGlzLDUpKzEscz1CKHRoaXMsNCkrNCx4PW5ldyAodD9VaW50OEFycmF5OkFycmF5KShDLmxlbmd0aCksUT1rLFI9ayxTPWssdj1rLE09ayxGPWssej1rLHE9ayxUPWsscT0wO3E8czsrK3EpeFtDW3FdXT0KQih0aGlzLDMpO2lmKCF0KXtxPXM7Zm9yKHM9eC5sZW5ndGg7cTxzOysrcSl4W0NbcV1dPTB9UT11KHgpO3Y9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKG0rcCk7cT0wO2ZvcihUPW0rcDtxPFQ7KXN3aXRjaChNPUQodGhpcyxRKSxNKXtjYXNlIDE2OmZvcih6PTMrQih0aGlzLDIpO3otLTspdltxKytdPUY7YnJlYWs7Y2FzZSAxNzpmb3Ioej0zK0IodGhpcywzKTt6LS07KXZbcSsrXT0wO0Y9MDticmVhaztjYXNlIDE4OmZvcih6PTExK0IodGhpcyw3KTt6LS07KXZbcSsrXT0wO0Y9MDticmVhaztkZWZhdWx0OkY9dltxKytdPU19Uj10P3Uodi5zdWJhcnJheSgwLG0pKTp1KHYuc2xpY2UoMCxtKSk7Uz10P3Uodi5zdWJhcnJheShtKSk6dSh2LnNsaWNlKG0pKTt0aGlzLmooUixTKTticmVhaztkZWZhdWx0OnRocm93IEVycm9yKCJ1bmtub3duIEJUWVBFOiAiK2MpO319cmV0dXJuIHRoaXMubSgpfTsKdmFyIEU9WzE2LDE3LDE4LDAsOCw3LDksNiwxMCw1LDExLDQsMTIsMywxMywyLDE0LDEsMTVdLEM9dD9uZXcgVWludDE2QXJyYXkoRSk6RSxHPVszLDQsNSw2LDcsOCw5LDEwLDExLDEzLDE1LDE3LDE5LDIzLDI3LDMxLDM1LDQzLDUxLDU5LDY3LDgzLDk5LDExNSwxMzEsMTYzLDE5NSwyMjcsMjU4LDI1OCwyNThdLEg9dD9uZXcgVWludDE2QXJyYXkoRyk6RyxJPVswLDAsMCwwLDAsMCwwLDAsMSwxLDEsMSwyLDIsMiwyLDMsMywzLDMsNCw0LDQsNCw1LDUsNSw1LDAsMCwwXSxKPXQ/bmV3IFVpbnQ4QXJyYXkoSSk6SSxLPVsxLDIsMyw0LDUsNyw5LDEzLDE3LDI1LDMzLDQ5LDY1LDk3LDEyOSwxOTMsMjU3LDM4NSw1MTMsNzY5LDEwMjUsMTUzNywyMDQ5LDMwNzMsNDA5Nyw2MTQ1LDgxOTMsMTIyODksMTYzODUsMjQ1NzddLEw9dD9uZXcgVWludDE2QXJyYXkoSyk6SyxOPVswLDAsMCwwLDEsMSwyLDIsMywzLDQsNCw1LDUsNiw2LDcsNyw4LDgsOSw5LDEwLDEwLDExLDExLDEyLDEyLDEzLAoxM10sTz10P25ldyBVaW50OEFycmF5KE4pOk4sUD1uZXcgKHQ/VWludDhBcnJheTpBcnJheSkoMjg4KSxVLGRhO1U9MDtmb3IoZGE9UC5sZW5ndGg7VTxkYTsrK1UpUFtVXT0xNDM+PVU/ODoyNTU+PVU/OToyNzk+PVU/Nzo4O3ZhciBiYT11KFApLFY9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKDMwKSxXLGVhO1c9MDtmb3IoZWE9Vi5sZW5ndGg7VzxlYTsrK1cpVltXXT01O3ZhciBjYT11KFYpO2Z1bmN0aW9uIEIoYyxkKXtmb3IodmFyIGE9Yy5mLGI9Yy5jLGU9Yy5pbnB1dCxmPWMuZCxnPWUubGVuZ3RoLGg7YjxkOyl7aWYoZj49Zyl0aHJvdyBFcnJvcigiaW5wdXQgYnVmZmVyIGlzIGJyb2tlbiIpO2F8PWVbZisrXTw8YjtiKz04fWg9YSYoMTw8ZCktMTtjLmY9YT4+PmQ7Yy5jPWItZDtjLmQ9ZjtyZXR1cm4gaH0KZnVuY3Rpb24gRChjLGQpe2Zvcih2YXIgYT1jLmYsYj1jLmMsZT1jLmlucHV0LGY9Yy5kLGc9ZS5sZW5ndGgsaD1kWzBdLGw9ZFsxXSxuLG07YjxsJiYhKGY+PWcpOylhfD1lW2YrK108PGIsYis9ODtuPWhbYSYoMTw8bCktMV07bT1uPj4+MTY7aWYobT5iKXRocm93IEVycm9yKCJpbnZhbGlkIGNvZGUgbGVuZ3RoOiAiK20pO2MuZj1hPj5tO2MuYz1iLW07Yy5kPWY7cmV0dXJuIG4mNjU1MzV9CncucHJvdG90eXBlLmo9ZnVuY3Rpb24oYyxkKXt2YXIgYT10aGlzLmIsYj10aGlzLmE7dGhpcy5uPWM7Zm9yKHZhciBlPWEubGVuZ3RoLTI1OCxmLGcsaCxsOzI1NiE9PShmPUQodGhpcyxjKSk7KWlmKDI1Nj5mKWI+PWUmJih0aGlzLmE9YixhPXRoaXMuZSgpLGI9dGhpcy5hKSxhW2IrK109ZjtlbHNle2c9Zi0yNTc7bD1IW2ddOzA8SltnXSYmKGwrPUIodGhpcyxKW2ddKSk7Zj1EKHRoaXMsZCk7aD1MW2ZdOzA8T1tmXSYmKGgrPUIodGhpcyxPW2ZdKSk7Yj49ZSYmKHRoaXMuYT1iLGE9dGhpcy5lKCksYj10aGlzLmEpO2Zvcig7bC0tOylhW2JdPWFbYisrLWhdfWZvcig7ODw9dGhpcy5jOyl0aGlzLmMtPTgsdGhpcy5kLS07dGhpcy5hPWJ9Owp3LnByb3RvdHlwZS5zPWZ1bmN0aW9uKGMsZCl7dmFyIGE9dGhpcy5iLGI9dGhpcy5hO3RoaXMubj1jO2Zvcih2YXIgZT1hLmxlbmd0aCxmLGcsaCxsOzI1NiE9PShmPUQodGhpcyxjKSk7KWlmKDI1Nj5mKWI+PWUmJihhPXRoaXMuZSgpLGU9YS5sZW5ndGgpLGFbYisrXT1mO2Vsc2V7Zz1mLTI1NztsPUhbZ107MDxKW2ddJiYobCs9Qih0aGlzLEpbZ10pKTtmPUQodGhpcyxkKTtoPUxbZl07MDxPW2ZdJiYoaCs9Qih0aGlzLE9bZl0pKTtiK2w+ZSYmKGE9dGhpcy5lKCksZT1hLmxlbmd0aCk7Zm9yKDtsLS07KWFbYl09YVtiKystaF19Zm9yKDs4PD10aGlzLmM7KXRoaXMuYy09OCx0aGlzLmQtLTt0aGlzLmE9Yn07CncucHJvdG90eXBlLmU9ZnVuY3Rpb24oKXt2YXIgYz1uZXcgKHQ/VWludDhBcnJheTpBcnJheSkodGhpcy5hLTMyNzY4KSxkPXRoaXMuYS0zMjc2OCxhLGIsZT10aGlzLmI7aWYodCljLnNldChlLnN1YmFycmF5KDMyNzY4LGMubGVuZ3RoKSk7ZWxzZXthPTA7Zm9yKGI9Yy5sZW5ndGg7YTxiOysrYSljW2FdPWVbYSszMjc2OF19dGhpcy5nLnB1c2goYyk7dGhpcy5rKz1jLmxlbmd0aDtpZih0KWUuc2V0KGUuc3ViYXJyYXkoZCxkKzMyNzY4KSk7ZWxzZSBmb3IoYT0wOzMyNzY4PmE7KythKWVbYV09ZVtkK2FdO3RoaXMuYT0zMjc2ODtyZXR1cm4gZX07CncucHJvdG90eXBlLnU9ZnVuY3Rpb24oYyl7dmFyIGQsYT10aGlzLmlucHV0Lmxlbmd0aC90aGlzLmQrMXwwLGIsZSxmLGc9dGhpcy5pbnB1dCxoPXRoaXMuYjtjJiYoIm51bWJlciI9PT10eXBlb2YgYy5vJiYoYT1jLm8pLCJudW1iZXIiPT09dHlwZW9mIGMucSYmKGErPWMucSkpOzI+YT8oYj0oZy5sZW5ndGgtdGhpcy5kKS90aGlzLm5bMl0sZj0yNTgqKGIvMil8MCxlPWY8aC5sZW5ndGg/aC5sZW5ndGgrZjpoLmxlbmd0aDw8MSk6ZT1oLmxlbmd0aCphO3Q/KGQ9bmV3IFVpbnQ4QXJyYXkoZSksZC5zZXQoaCkpOmQ9aDtyZXR1cm4gdGhpcy5iPWR9Owp3LnByb3RvdHlwZS5tPWZ1bmN0aW9uKCl7dmFyIGM9MCxkPXRoaXMuYixhPXRoaXMuZyxiLGU9bmV3ICh0P1VpbnQ4QXJyYXk6QXJyYXkpKHRoaXMuaysodGhpcy5hLTMyNzY4KSksZixnLGgsbDtpZigwPT09YS5sZW5ndGgpcmV0dXJuIHQ/dGhpcy5iLnN1YmFycmF5KDMyNzY4LHRoaXMuYSk6dGhpcy5iLnNsaWNlKDMyNzY4LHRoaXMuYSk7Zj0wO2ZvcihnPWEubGVuZ3RoO2Y8ZzsrK2Ype2I9YVtmXTtoPTA7Zm9yKGw9Yi5sZW5ndGg7aDxsOysraCllW2MrK109YltoXX1mPTMyNzY4O2ZvcihnPXRoaXMuYTtmPGc7KytmKWVbYysrXT1kW2ZdO3RoaXMuZz1bXTtyZXR1cm4gdGhpcy5idWZmZXI9ZX07CncucHJvdG90eXBlLnI9ZnVuY3Rpb24oKXt2YXIgYyxkPXRoaXMuYTt0P3RoaXMucD8oYz1uZXcgVWludDhBcnJheShkKSxjLnNldCh0aGlzLmIuc3ViYXJyYXkoMCxkKSkpOmM9dGhpcy5iLnN1YmFycmF5KDAsZCk6KHRoaXMuYi5sZW5ndGg+ZCYmKHRoaXMuYi5sZW5ndGg9ZCksYz10aGlzLmIpO3JldHVybiB0aGlzLmJ1ZmZlcj1jfTtyKCJabGliLlJhd0luZmxhdGUiLHcpO3IoIlpsaWIuUmF3SW5mbGF0ZS5wcm90b3R5cGUuZGVjb21wcmVzcyIsdy5wcm90b3R5cGUudCk7dmFyIFg9e0FEQVBUSVZFOnksQkxPQ0s6QX0sWSxaLCQsZmE7aWYoT2JqZWN0LmtleXMpWT1PYmplY3Qua2V5cyhYKTtlbHNlIGZvcihaIGluIFk9W10sJD0wLFgpWVskKytdPVo7JD0wO2ZvcihmYT1ZLmxlbmd0aDskPGZhOysrJClaPVlbJF0scigiWmxpYi5SYXdJbmZsYXRlLkJ1ZmZlclR5cGUuIitaLFhbWl0pO30pLmNhbGwodGhpcyk7Cg=="

var bootstrapCode = "Ly8gQ29weXJpZ2h0IDIwMTggVGhlIEdvIEF1dGhvcnMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCi8vIFVzZSBvZiB0aGlzIHNvdXJjZSBjb2RlIGlzIGdvdmVybmVkIGJ5IGEgQlNELXN0eWxlCi8vIGxpY2Vuc2UgdGhhdCBjYW4gYmUgZm91bmQgaW4gdGhlIExJQ0VOU0UgZmlsZS4KCigoKSA9PiB7CgkvLyBNYXAgd2ViIGJyb3dzZXIgQVBJIGFuZCBOb2RlLmpzIEFQSSB0byBhIHNpbmdsZSBjb21tb24gQVBJIChwcmVmZXJyaW5nIHdlYiBzdGFuZGFyZHMgb3ZlciBOb2RlLmpzIEFQSSkuCgljb25zdCBpc05vZGVKUyA9IHR5cGVvZiBwcm9jZXNzICE9PSAidW5kZWZpbmVkIjsKCWlmIChpc05vZGVKUykgewoJCWNvbnN0IG5vZGVDcnlwdG8gPSByZXF1aXJlKCJjcnlwdG8iKTsKCQlnbG9iYWwuY3J5cHRvID0gewoJCQlnZXRSYW5kb21WYWx1ZXMoYikgewoJCQkJbm9kZUNyeXB0by5yYW5kb21GaWxsU3luYyhiKTsKCQkJfSwKCQl9OwoKCQlnbG9iYWwucGVyZm9ybWFuY2UgPSB7CgkJCW5vdygpIHsKCQkJCWNvbnN0IFtzZWMsIG5zZWNdID0gcHJvY2Vzcy5ocnRpbWUoKTsKCQkJCXJldHVybiBzZWMgKiAxMDAwICsgbnNlYyAvIDEwMDAwMDA7CgkJCX0sCgkJfTsKCgkJY29uc3QgdXRpbCA9IHJlcXVpcmUoInV0aWwiKTsKCQlnbG9iYWwuVGV4dEVuY29kZXIgPSB1dGlsLlRleHRFbmNvZGVyOwoJCWdsb2JhbC5UZXh0RGVjb2RlciA9IHV0aWwuVGV4dERlY29kZXI7CgkJZ2xvYmFsLndpbmRvdyA9IGdsb2JhbDsKCX0gCgoJaWYgKHR5cGVvZiB3aW5kb3cgIT09ICJ1bmRlZmluZWQiKSB7CgkJd2luZG93Lmdsb2JhbCA9IHdpbmRvdzsKCX0gZWxzZSBpZiAodHlwZW9mIHNlbGYgIT09ICJ1bmRlZmluZWQiKSB7CgkJc2VsZi5nbG9iYWwgPSBzZWxmOwoJfSBlbHNlIHsKCQl0aHJvdyBuZXcgRXJyb3IoImNhbm5vdCBleHBvcnQgR28gKG5laXRoZXIgd2luZG93IG5vciBzZWxmIGlzIGRlZmluZWQpIik7Cgl9CgoJbGV0IG91dHB1dEJ1ZiA9ICIiOwoJZ2xvYmFsLmZzID0gewoJCWNvbnN0YW50czogeyBPX1dST05MWTogLTEsIE9fUkRXUjogLTEsIE9fQ1JFQVQ6IC0xLCBPX1RSVU5DOiAtMSwgT19BUFBFTkQ6IC0xLCBPX0VYQ0w6IC0xIH0sIC8vIHVudXNlZAoJCXdyaXRlU3luYyhmZCwgYnVmKSB7CgkJCW91dHB1dEJ1ZiArPSBkZWNvZGVyLmRlY29kZShidWYpOwoJCQljb25zdCBubCA9IG91dHB1dEJ1Zi5sYXN0SW5kZXhPZigiXG4iKTsKCQkJaWYgKG5sICE9IC0xKSB7CgkJCQljb25zb2xlLmxvZyhvdXRwdXRCdWYuc3Vic3RyKDAsIG5sKSk7CgkJCQlvdXRwdXRCdWYgPSBvdXRwdXRCdWYuc3Vic3RyKG5sICsgMSk7CgkJCX0KCQkJcmV0dXJuIGJ1Zi5sZW5ndGg7CgkJfSwKCQlvcGVuU3luYyhwYXRoLCBmbGFncywgbW9kZSkgewoJCQljb25zdCBlcnIgPSBuZXcgRXJyb3IoIm5vdCBpbXBsZW1lbnRlZCIpOwoJCQllcnIuY29kZSA9ICJFTk9TWVMiOwoJCQl0aHJvdyBlcnI7CgkJfSwKCX07CgoKCWNvbnN0IGVuY29kZXIgPSBuZXcgVGV4dEVuY29kZXIoInV0Zi04Iik7Cgljb25zdCBkZWNvZGVyID0gbmV3IFRleHREZWNvZGVyKCJ1dGYtOCIpOwoKCWdsb2JhbC5HbyA9IGNsYXNzIHsKCQljb25zdHJ1Y3RvcigpIHsKCQkJdGhpcy5hcmd2ID0gWyJqcyJdOwoJCQl0aGlzLmVudiA9IHt9OwoJCQl0aGlzLmV4aXQgPSAoY29kZSkgPT4gewoJCQkJaWYgKGNvZGUgIT09IDApIHsKCQkJCQljb25zb2xlLndhcm4oImV4aXQgY29kZToiLCBjb2RlKTsKCQkJCX0KCQkJfTsKCQkJdGhpcy5fY2FsbGJhY2tUaW1lb3V0cyA9IG5ldyBNYXAoKTsKCQkJdGhpcy5fbmV4dENhbGxiYWNrVGltZW91dElEID0gMTsKCgkJCWNvbnN0IG1lbSA9ICgpID0+IHsKCQkJCS8vIFRoZSBidWZmZXIgbWF5IGNoYW5nZSB3aGVuIHJlcXVlc3RpbmcgbW9yZSBtZW1vcnkuCgkJCQlyZXR1cm4gbmV3IERhdGFWaWV3KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyKTsKCQkJfQoKCQkJY29uc3Qgc2V0SW50NjQgPSAoYWRkciwgdikgPT4gewoJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIgKyAwLCB2LCB0cnVlKTsKCQkJCW1lbSgpLnNldFVpbnQzMihhZGRyICsgNCwgTWF0aC5mbG9vcih2IC8gNDI5NDk2NzI5NiksIHRydWUpOwoJCQl9CgoJCQljb25zdCBnZXRJbnQ2NCA9IChhZGRyKSA9PiB7CgkJCQljb25zdCBsb3cgPSBtZW0oKS5nZXRVaW50MzIoYWRkciArIDAsIHRydWUpOwoJCQkJY29uc3QgaGlnaCA9IG1lbSgpLmdldEludDMyKGFkZHIgKyA0LCB0cnVlKTsKCQkJCXJldHVybiBsb3cgKyBoaWdoICogNDI5NDk2NzI5NjsKCQkJfQoKCQkJY29uc3QgbG9hZFZhbHVlID0gKGFkZHIpID0+IHsKCQkJCWNvbnN0IGYgPSBtZW0oKS5nZXRGbG9hdDY0KGFkZHIsIHRydWUpOwoJCQkJaWYgKCFpc05hTihmKSkgewoJCQkJCXJldHVybiBmOwoJCQkJfQoKCQkJCWNvbnN0IGlkID0gbWVtKCkuZ2V0VWludDMyKGFkZHIsIHRydWUpOwoJCQkJcmV0dXJuIHRoaXMuX3ZhbHVlc1tpZF07CgkJCX0KCgkJCWNvbnN0IHN0b3JlVmFsdWUgPSAoYWRkciwgdikgPT4gewoJCQkJY29uc3QgbmFuSGVhZCA9IDB4N0ZGODAwMDA7CgoJCQkJaWYgKHR5cGVvZiB2ID09PSAibnVtYmVyIikgewoJCQkJCWlmIChpc05hTih2KSkgewoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQsIHRydWUpOwoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciwgMCwgdHJ1ZSk7CgkJCQkJCXJldHVybjsKCQkJCQl9CgkJCQkJbWVtKCkuc2V0RmxvYXQ2NChhZGRyLCB2LCB0cnVlKTsKCQkJCQlyZXR1cm47CgkJCQl9CgoJCQkJc3dpdGNoICh2KSB7CgkJCQkJY2FzZSB1bmRlZmluZWQ6CgkJCQkJCW1lbSgpLnNldFVpbnQzMihhZGRyICsgNCwgbmFuSGVhZCwgdHJ1ZSk7CgkJCQkJCW1lbSgpLnNldFVpbnQzMihhZGRyLCAxLCB0cnVlKTsKCQkJCQkJcmV0dXJuOwoJCQkJCWNhc2UgbnVsbDoKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIgKyA0LCBuYW5IZWFkLCB0cnVlKTsKCQkJCQkJbWVtKCkuc2V0VWludDMyKGFkZHIsIDIsIHRydWUpOwoJCQkJCQlyZXR1cm47CgkJCQkJY2FzZSB0cnVlOgoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQsIHRydWUpOwoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciwgMywgdHJ1ZSk7CgkJCQkJCXJldHVybjsKCQkJCQljYXNlIGZhbHNlOgoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQsIHRydWUpOwoJCQkJCQltZW0oKS5zZXRVaW50MzIoYWRkciwgNCwgdHJ1ZSk7CgkJCQkJCXJldHVybjsKCQkJCX0KCgkJCQlsZXQgcmVmID0gdGhpcy5fcmVmcy5nZXQodik7CgkJCQlpZiAocmVmID09PSB1bmRlZmluZWQpIHsKCQkJCQlyZWYgPSB0aGlzLl92YWx1ZXMubGVuZ3RoOwoJCQkJCXRoaXMuX3ZhbHVlcy5wdXNoKHYpOwoJCQkJCXRoaXMuX3JlZnMuc2V0KHYsIHJlZik7CgkJCQl9CgkJCQlsZXQgdHlwZUZsYWcgPSAwOwoJCQkJc3dpdGNoICh0eXBlb2YgdikgewoJCQkJCWNhc2UgInN0cmluZyI6CgkJCQkJCXR5cGVGbGFnID0gMTsKCQkJCQkJYnJlYWs7CgkJCQkJY2FzZSAic3ltYm9sIjoKCQkJCQkJdHlwZUZsYWcgPSAyOwoJCQkJCQlicmVhazsKCQkJCQljYXNlICJmdW5jdGlvbiI6CgkJCQkJCXR5cGVGbGFnID0gMzsKCQkJCQkJYnJlYWs7CgkJCQl9CgkJCQltZW0oKS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQgfCB0eXBlRmxhZywgdHJ1ZSk7CgkJCQltZW0oKS5zZXRVaW50MzIoYWRkciwgcmVmLCB0cnVlKTsKCQkJfQoKCQkJY29uc3QgbG9hZFNsaWNlID0gKGFkZHIpID0+IHsKCQkJCWNvbnN0IGFycmF5ID0gZ2V0SW50NjQoYWRkciArIDApOwoJCQkJY29uc3QgbGVuID0gZ2V0SW50NjQoYWRkciArIDgpOwoJCQkJcmV0dXJuIG5ldyBVaW50OEFycmF5KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyLCBhcnJheSwgbGVuKTsKCQkJfQoKCQkJY29uc3QgbG9hZFNsaWNlT2ZWYWx1ZXMgPSAoYWRkcikgPT4gewoJCQkJY29uc3QgYXJyYXkgPSBnZXRJbnQ2NChhZGRyICsgMCk7CgkJCQljb25zdCBsZW4gPSBnZXRJbnQ2NChhZGRyICsgOCk7CgkJCQljb25zdCBhID0gbmV3IEFycmF5KGxlbik7CgkJCQlmb3IgKGxldCBpID0gMDsgaSA8IGxlbjsgaSsrKSB7CgkJCQkJYVtpXSA9IGxvYWRWYWx1ZShhcnJheSArIGkgKiA4KTsKCQkJCX0KCQkJCXJldHVybiBhOwoJCQl9CgoJCQljb25zdCBsb2FkU3RyaW5nID0gKGFkZHIpID0+IHsKCQkJCWNvbnN0IHNhZGRyID0gZ2V0SW50NjQoYWRkciArIDApOwoJCQkJY29uc3QgbGVuID0gZ2V0SW50NjQoYWRkciArIDgpOwoJCQkJcmV0dXJuIGRlY29kZXIuZGVjb2RlKG5ldyBEYXRhVmlldyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlciwgc2FkZHIsIGxlbikpOwoJCQl9CgoJCQljb25zdCB0aW1lT3JpZ2luID0gRGF0ZS5ub3coKSAtIHBlcmZvcm1hbmNlLm5vdygpOwoJCQl0aGlzLmltcG9ydE9iamVjdCA9IHsKCQkJCWdvOiB7CgkJCQkJLy8gZnVuYyB3YXNtRXhpdChjb2RlIGludDMyKQoJCQkJCSJydW50aW1lLndhc21FeGl0IjogKHNwKSA9PiB7CgkJCQkJCWNvbnN0IGNvZGUgPSBtZW0oKS5nZXRJbnQzMihzcCArIDgsIHRydWUpOwoJCQkJCQl0aGlzLmV4aXRlZCA9IHRydWU7CgkJCQkJCWRlbGV0ZSB0aGlzLl9pbnN0OwoJCQkJCQlkZWxldGUgdGhpcy5fdmFsdWVzOwoJCQkJCQlkZWxldGUgdGhpcy5fcmVmczsKCQkJCQkJdGhpcy5leGl0KGNvZGUpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgd2FzbVdyaXRlKGZkIHVpbnRwdHIsIHAgdW5zYWZlLlBvaW50ZXIsIG4gaW50MzIpCgkJCQkJInJ1bnRpbWUud2FzbVdyaXRlIjogKHNwKSA9PiB7CgkJCQkJCWNvbnN0IGZkID0gZ2V0SW50NjQoc3AgKyA4KTsKCQkJCQkJY29uc3QgcCA9IGdldEludDY0KHNwICsgMTYpOwoJCQkJCQljb25zdCBuID0gbWVtKCkuZ2V0SW50MzIoc3AgKyAyNCwgdHJ1ZSk7CgkJCQkJCWZzLndyaXRlU3luYyhmZCwgbmV3IFVpbnQ4QXJyYXkodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIsIHAsIG4pKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIG5hbm90aW1lKCkgaW50NjQKCQkJCQkicnVudGltZS5uYW5vdGltZSI6IChzcCkgPT4gewoJCQkJCQlzZXRJbnQ2NChzcCArIDgsICh0aW1lT3JpZ2luICsgcGVyZm9ybWFuY2Uubm93KCkpICogMTAwMDAwMCk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB3YWxsdGltZSgpIChzZWMgaW50NjQsIG5zZWMgaW50MzIpCgkJCQkJInJ1bnRpbWUud2FsbHRpbWUiOiAoc3ApID0+IHsKCQkJCQkJY29uc3QgbXNlYyA9IChuZXcgRGF0ZSkuZ2V0VGltZSgpOwoJCQkJCQlzZXRJbnQ2NChzcCArIDgsIG1zZWMgLyAxMDAwKTsKCQkJCQkJbWVtKCkuc2V0SW50MzIoc3AgKyAxNiwgKG1zZWMgJSAxMDAwKSAqIDEwMDAwMDAsIHRydWUpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgc2NoZWR1bGVDYWxsYmFjayhkZWxheSBpbnQ2NCkgaW50MzIKCQkJCQkicnVudGltZS5zY2hlZHVsZUNhbGxiYWNrIjogKHNwKSA9PiB7CgkJCQkJCWNvbnN0IGlkID0gdGhpcy5fbmV4dENhbGxiYWNrVGltZW91dElEOwoJCQkJCQl0aGlzLl9uZXh0Q2FsbGJhY2tUaW1lb3V0SUQrKzsKCQkJCQkJdGhpcy5fY2FsbGJhY2tUaW1lb3V0cy5zZXQoaWQsIHNldFRpbWVvdXQoCgkJCQkJCQkoKSA9PiB7IHRoaXMuX3Jlc29sdmVDYWxsYmFja1Byb21pc2UoKTsgfSwKCQkJCQkJCWdldEludDY0KHNwICsgOCkgKyAxLCAvLyBzZXRUaW1lb3V0IGhhcyBiZWVuIHNlZW4gdG8gZmlyZSB1cCB0byAxIG1pbGxpc2Vjb25kIGVhcmx5CgkJCQkJCSkpOwoJCQkJCQltZW0oKS5zZXRJbnQzMihzcCArIDE2LCBpZCwgdHJ1ZSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyBjbGVhclNjaGVkdWxlZENhbGxiYWNrKGlkIGludDMyKQoJCQkJCSJydW50aW1lLmNsZWFyU2NoZWR1bGVkQ2FsbGJhY2siOiAoc3ApID0+IHsKCQkJCQkJY29uc3QgaWQgPSBtZW0oKS5nZXRJbnQzMihzcCArIDgsIHRydWUpOwoJCQkJCQljbGVhclRpbWVvdXQodGhpcy5fY2FsbGJhY2tUaW1lb3V0cy5nZXQoaWQpKTsKCQkJCQkJdGhpcy5fY2FsbGJhY2tUaW1lb3V0cy5kZWxldGUoaWQpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgZ2V0UmFuZG9tRGF0YShyIFtdYnl0ZSkKCQkJCQkicnVudGltZS5nZXRSYW5kb21EYXRhIjogKHNwKSA9PiB7CgkJCQkJCWNyeXB0by5nZXRSYW5kb21WYWx1ZXMobG9hZFNsaWNlKHNwICsgOCkpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgc3RyaW5nVmFsKHZhbHVlIHN0cmluZykgcmVmCgkJCQkJInN5c2NhbGwvanMuc3RyaW5nVmFsIjogKHNwKSA9PiB7CgkJCQkJCXN0b3JlVmFsdWUoc3AgKyAyNCwgbG9hZFN0cmluZyhzcCArIDgpKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHZhbHVlR2V0KHYgcmVmLCBwIHN0cmluZykgcmVmCgkJCQkJInN5c2NhbGwvanMudmFsdWVHZXQiOiAoc3ApID0+IHsKCQkJCQkJc3RvcmVWYWx1ZShzcCArIDMyLCBSZWZsZWN0LmdldChsb2FkVmFsdWUoc3AgKyA4KSwgbG9hZFN0cmluZyhzcCArIDE2KSkpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVTZXQodiByZWYsIHAgc3RyaW5nLCB4IHJlZikKCQkJCQkic3lzY2FsbC9qcy52YWx1ZVNldCI6IChzcCkgPT4gewoJCQkJCQlSZWZsZWN0LnNldChsb2FkVmFsdWUoc3AgKyA4KSwgbG9hZFN0cmluZyhzcCArIDE2KSwgbG9hZFZhbHVlKHNwICsgMzIpKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHZhbHVlSW5kZXgodiByZWYsIGkgaW50KSByZWYKCQkJCQkic3lzY2FsbC9qcy52YWx1ZUluZGV4IjogKHNwKSA9PiB7CgkJCQkJCXN0b3JlVmFsdWUoc3AgKyAyNCwgUmVmbGVjdC5nZXQobG9hZFZhbHVlKHNwICsgOCksIGdldEludDY0KHNwICsgMTYpKSk7CgkJCQkJfSwKCgkJCQkJLy8gdmFsdWVTZXRJbmRleCh2IHJlZiwgaSBpbnQsIHggcmVmKQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlU2V0SW5kZXgiOiAoc3ApID0+IHsKCQkJCQkJUmVmbGVjdC5zZXQobG9hZFZhbHVlKHNwICsgOCksIGdldEludDY0KHNwICsgMTYpLCBsb2FkVmFsdWUoc3AgKyAyNCkpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVDYWxsKHYgcmVmLCBtIHN0cmluZywgYXJncyBbXXJlZikgKHJlZiwgYm9vbCkKCQkJCQkic3lzY2FsbC9qcy52YWx1ZUNhbGwiOiAoc3ApID0+IHsKCQkJCQkJdHJ5IHsKCQkJCQkJCWNvbnN0IHYgPSBsb2FkVmFsdWUoc3AgKyA4KTsKCQkJCQkJCWNvbnN0IG0gPSBSZWZsZWN0LmdldCh2LCBsb2FkU3RyaW5nKHNwICsgMTYpKTsKCQkJCQkJCWNvbnN0IGFyZ3MgPSBsb2FkU2xpY2VPZlZhbHVlcyhzcCArIDMyKTsKCQkJCQkJCXN0b3JlVmFsdWUoc3AgKyA1NiwgUmVmbGVjdC5hcHBseShtLCB2LCBhcmdzKSk7CgkJCQkJCQltZW0oKS5zZXRVaW50OChzcCArIDY0LCAxKTsKCQkJCQkJfSBjYXRjaCAoZXJyKSB7CgkJCQkJCQlzdG9yZVZhbHVlKHNwICsgNTYsIGVycik7CgkJCQkJCQltZW0oKS5zZXRVaW50OChzcCArIDY0LCAwKTsKCQkJCQkJfQoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVJbnZva2UodiByZWYsIGFyZ3MgW11yZWYpIChyZWYsIGJvb2wpCgkJCQkJInN5c2NhbGwvanMudmFsdWVJbnZva2UiOiAoc3ApID0+IHsKCQkJCQkJdHJ5IHsKCQkJCQkJCWNvbnN0IHYgPSBsb2FkVmFsdWUoc3AgKyA4KTsKCQkJCQkJCWNvbnN0IGFyZ3MgPSBsb2FkU2xpY2VPZlZhbHVlcyhzcCArIDE2KTsKCQkJCQkJCXN0b3JlVmFsdWUoc3AgKyA0MCwgUmVmbGVjdC5hcHBseSh2LCB1bmRlZmluZWQsIGFyZ3MpKTsKCQkJCQkJCW1lbSgpLnNldFVpbnQ4KHNwICsgNDgsIDEpOwoJCQkJCQl9IGNhdGNoIChlcnIpIHsKCQkJCQkJCXN0b3JlVmFsdWUoc3AgKyA0MCwgZXJyKTsKCQkJCQkJCW1lbSgpLnNldFVpbnQ4KHNwICsgNDgsIDApOwoJCQkJCQl9CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB2YWx1ZU5ldyh2IHJlZiwgYXJncyBbXXJlZikgKHJlZiwgYm9vbCkKCQkJCQkic3lzY2FsbC9qcy52YWx1ZU5ldyI6IChzcCkgPT4gewoJCQkJCQl0cnkgewoJCQkJCQkJY29uc3QgdiA9IGxvYWRWYWx1ZShzcCArIDgpOwoJCQkJCQkJY29uc3QgYXJncyA9IGxvYWRTbGljZU9mVmFsdWVzKHNwICsgMTYpOwoJCQkJCQkJc3RvcmVWYWx1ZShzcCArIDQwLCBSZWZsZWN0LmNvbnN0cnVjdCh2LCBhcmdzKSk7CgkJCQkJCQltZW0oKS5zZXRVaW50OChzcCArIDQ4LCAxKTsKCQkJCQkJfSBjYXRjaCAoZXJyKSB7CgkJCQkJCQlzdG9yZVZhbHVlKHNwICsgNDAsIGVycik7CgkJCQkJCQltZW0oKS5zZXRVaW50OChzcCArIDQ4LCAwKTsKCQkJCQkJfQoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVMZW5ndGgodiByZWYpIGludAoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlTGVuZ3RoIjogKHNwKSA9PiB7CgkJCQkJCXNldEludDY0KHNwICsgMTYsIHBhcnNlSW50KGxvYWRWYWx1ZShzcCArIDgpLmxlbmd0aCkpOwoJCQkJCX0sCgoJCQkJCS8vIHZhbHVlUHJlcGFyZVN0cmluZyh2IHJlZikgKHJlZiwgaW50KQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlUHJlcGFyZVN0cmluZyI6IChzcCkgPT4gewoJCQkJCQljb25zdCBzdHIgPSBlbmNvZGVyLmVuY29kZShTdHJpbmcobG9hZFZhbHVlKHNwICsgOCkpKTsKCQkJCQkJc3RvcmVWYWx1ZShzcCArIDE2LCBzdHIpOwoJCQkJCQlzZXRJbnQ2NChzcCArIDI0LCBzdHIubGVuZ3RoKTsKCQkJCQl9LAoKCQkJCQkvLyB2YWx1ZUxvYWRTdHJpbmcodiByZWYsIGIgW11ieXRlKQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlTG9hZFN0cmluZyI6IChzcCkgPT4gewoJCQkJCQljb25zdCBzdHIgPSBsb2FkVmFsdWUoc3AgKyA4KTsKCQkJCQkJbG9hZFNsaWNlKHNwICsgMTYpLnNldChzdHIpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVJbnN0YW5jZU9mKHYgcmVmLCB0IHJlZikgYm9vbAoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlSW5zdGFuY2VPZiI6IChzcCkgPT4gewoJCQkJCQltZW0oKS5zZXRVaW50OChzcCArIDI0LCBsb2FkVmFsdWUoc3AgKyA4KSBpbnN0YW5jZW9mIGxvYWRWYWx1ZShzcCArIDE2KSk7CgkJCQkJfSwKCgkJCQkJImRlYnVnIjogKHZhbHVlKSA9PiB7CgkJCQkJCWNvbnNvbGUubG9nKHZhbHVlKTsKCQkJCQl9LAoJCQkJfQoJCQl9OwoJCX0KCgkJYXN5bmMgcnVuKGluc3RhbmNlKSB7CgkJCXRoaXMuX2luc3QgPSBpbnN0YW5jZTsKCQkJdGhpcy5fdmFsdWVzID0gWyAvLyBUT0RPOiBnYXJiYWdlIGNvbGxlY3Rpb24KCQkJCU5hTiwKCQkJCXVuZGVmaW5lZCwKCQkJCW51bGwsCgkJCQl0cnVlLAoJCQkJZmFsc2UsCgkJCQlnbG9iYWwsCgkJCQl0aGlzLl9pbnN0LmV4cG9ydHMubWVtLAoJCQkJdGhpcywKCQkJXTsKCQkJdGhpcy5fcmVmcyA9IG5ldyBNYXAoKTsKCQkJdGhpcy5fY2FsbGJhY2tTaHV0ZG93biA9IGZhbHNlOwoJCQl0aGlzLmV4aXRlZCA9IGZhbHNlOwoKCQkJY29uc3QgbWVtID0gbmV3IERhdGFWaWV3KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyKQoKCQkJLy8gUGFzcyBjb21tYW5kIGxpbmUgYXJndW1lbnRzIGFuZCBlbnZpcm9ubWVudCB2YXJpYWJsZXMgdG8gV2ViQXNzZW1ibHkgYnkgd3JpdGluZyB0aGVtIHRvIHRoZSBsaW5lYXIgbWVtb3J5LgoJCQlsZXQgb2Zmc2V0ID0gNDA5NjsKCgkJCWNvbnN0IHN0clB0ciA9IChzdHIpID0+IHsKCQkJCWxldCBwdHIgPSBvZmZzZXQ7CgkJCQluZXcgVWludDhBcnJheShtZW0uYnVmZmVyLCBvZmZzZXQsIHN0ci5sZW5ndGggKyAxKS5zZXQoZW5jb2Rlci5lbmNvZGUoc3RyICsgIlwwIikpOwoJCQkJb2Zmc2V0ICs9IHN0ci5sZW5ndGggKyAoOCAtIChzdHIubGVuZ3RoICUgOCkpOwoJCQkJcmV0dXJuIHB0cjsKCQkJfTsKCgkJCWNvbnN0IGFyZ2MgPSB0aGlzLmFyZ3YubGVuZ3RoOwoKCQkJY29uc3QgYXJndlB0cnMgPSBbXTsKCQkJdGhpcy5hcmd2LmZvckVhY2goKGFyZykgPT4gewoJCQkJYXJndlB0cnMucHVzaChzdHJQdHIoYXJnKSk7CgkJCX0pOwoKCQkJY29uc3Qga2V5cyA9IE9iamVjdC5rZXlzKHRoaXMuZW52KS5zb3J0KCk7CgkJCWFyZ3ZQdHJzLnB1c2goa2V5cy5sZW5ndGgpOwoJCQlrZXlzLmZvckVhY2goKGtleSkgPT4gewoJCQkJYXJndlB0cnMucHVzaChzdHJQdHIoYCR7a2V5fT0ke3RoaXMuZW52W2tleV19YCkpOwoJCQl9KTsKCgkJCWNvbnN0IGFyZ3YgPSBvZmZzZXQ7CgkJCWFyZ3ZQdHJzLmZvckVhY2goKHB0cikgPT4gewoJCQkJbWVtLnNldFVpbnQzMihvZmZzZXQsIHB0ciwgdHJ1ZSk7CgkJCQltZW0uc2V0VWludDMyKG9mZnNldCArIDQsIDAsIHRydWUpOwoJCQkJb2Zmc2V0ICs9IDg7CgkJCX0pOwoKCQkJd2hpbGUgKHRydWUpIHsKCQkJCWNvbnN0IGNhbGxiYWNrUHJvbWlzZSA9IG5ldyBQcm9taXNlKChyZXNvbHZlKSA9PiB7CgkJCQkJdGhpcy5fcmVzb2x2ZUNhbGxiYWNrUHJvbWlzZSA9ICgpID0+IHsKCQkJCQkJaWYgKHRoaXMuZXhpdGVkKSB7CgkJCQkJCQl0aHJvdyBuZXcgRXJyb3IoImJhZCBjYWxsYmFjazogR28gcHJvZ3JhbSBoYXMgYWxyZWFkeSBleGl0ZWQiKTsKCQkJCQkJfQoJCQkJCQlzZXRUaW1lb3V0KHJlc29sdmUsIDApOyAvLyBtYWtlIHN1cmUgaXQgaXMgYXN5bmNocm9ub3VzCgkJCQkJfTsKCQkJCX0pOwoJCQkJdGhpcy5faW5zdC5leHBvcnRzLnJ1bihhcmdjLCBhcmd2KTsKCQkJCWlmICh0aGlzLmV4aXRlZCkgewoJCQkJCWJyZWFrOwoJCQkJfQoJCQkJYXdhaXQgY2FsbGJhY2tQcm9taXNlOwoJCQl9CgkJfQoKCQlzdGF0aWMgX21ha2VDYWxsYmFja0hlbHBlcihpZCwgcGVuZGluZ0NhbGxiYWNrcywgZ28pIHsKCQkJcmV0dXJuIGZ1bmN0aW9uKCkgewoJCQkJcGVuZGluZ0NhbGxiYWNrcy5wdXNoKHsgaWQ6IGlkLCBhcmdzOiBhcmd1bWVudHMgfSk7CgkJCQlnby5fcmVzb2x2ZUNhbGxiYWNrUHJvbWlzZSgpOwoJCQl9OwoJCX0KCgkJc3RhdGljIF9tYWtlRXZlbnRDYWxsYmFja0hlbHBlcihwcmV2ZW50RGVmYXVsdCwgc3RvcFByb3BhZ2F0aW9uLCBzdG9wSW1tZWRpYXRlUHJvcGFnYXRpb24sIGZuKSB7CgkJCXJldHVybiBmdW5jdGlvbihldmVudCkgewoJCQkJaWYgKHByZXZlbnREZWZhdWx0KSB7CgkJCQkJZXZlbnQucHJldmVudERlZmF1bHQoKTsKCQkJCX0KCQkJCWlmIChzdG9wUHJvcGFnYXRpb24pIHsKCQkJCQlldmVudC5zdG9wUHJvcGFnYXRpb24oKTsKCQkJCX0KCQkJCWlmIChzdG9wSW1tZWRpYXRlUHJvcGFnYXRpb24pIHsKCQkJCQlldmVudC5zdG9wSW1tZWRpYXRlUHJvcGFnYXRpb24oKTsKCQkJCX0KCQkJCWZuKGV2ZW50KTsKCQkJfTsKCQl9Cgl9Cgp9KSgpOwo="
