"".init.1 t=1 size=26 args=0x0 locals=0x0
	0x0000 00000 (test9.go:10)	TEXT	"".init.1(SB), $0-0
	0x0000 00000 (test9.go:10)	MOVL	TLS, CX
	0x0007 00007 (test9.go:10)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:10)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:10)	JLS	19
	0x0012 00018 (test9.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (test9.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (test9.go:12)	RET
	0x0013 00019 (test9.go:12)	NOP
	0x0013 00019 (test9.go:10)	PCDATA	$0, $-1
	0x0013 00019 (test9.go:10)	CALL	runtime.morestack_noctxt(SB)
	0x0018 00024 (test9.go:10)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 76 01 c3 e8 00 00 00 00 eb e6                    v.........
	rel 9+4 t=16 TLS+0
	rel 20+4 t=8 runtime.morestack_noctxt+0
"".main t=1 size=632 args=0x0 locals=0x50
	0x0000 00000 (test9.go:24)	TEXT	"".main(SB), $80-0
	0x0000 00000 (test9.go:24)	MOVL	TLS, CX
	0x0007 00007 (test9.go:24)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:24)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:24)	JLS	622
	0x0016 00022 (test9.go:24)	SUBL	$80, SP
	0x0019 00025 (test9.go:24)	FUNCDATA	$0, gclocals·0ce64bbc7cfa5ef04d41c861de81a3d7(SB)
	0x0019 00025 (test9.go:24)	FUNCDATA	$1, gclocals·885e690be6afa9a2a3e0e4279d6dbcbc(SB)
	0x0019 00025 (test9.go:25)	LEAL	type."".Human(SB), AX
	0x001f 00031 (test9.go:25)	MOVL	AX, (SP)
	0x0022 00034 (test9.go:25)	PCDATA	$0, $0
	0x0022 00034 (test9.go:25)	CALL	runtime.newobject(SB)
	0x0027 00039 (test9.go:25)	MOVL	4(SP), AX
	0x002b 00043 (test9.go:25)	MOVL	AX, "".h+24(SP)
	0x002f 00047 (test9.go:26)	MOVL	$3, 4(AX)
	0x0036 00054 (test9.go:26)	MOVL	runtime.writeBarrier(SB), CX
	0x003c 00060 (test9.go:26)	TESTL	CX, CX
	0x003e 00062 (test9.go:26)	JNE	$0, 595
	0x0044 00068 (test9.go:26)	LEAL	go.string."Ben"(SB), CX
	0x004a 00074 (test9.go:26)	MOVL	CX, (AX)
	0x004c 00076 (test9.go:27)	LEAL	type."".Dog(SB), CX
	0x0052 00082 (test9.go:27)	MOVL	CX, (SP)
	0x0055 00085 (test9.go:27)	PCDATA	$0, $1
	0x0055 00085 (test9.go:27)	CALL	runtime.newobject(SB)
	0x005a 00090 (test9.go:27)	MOVL	4(SP), AX
	0x005e 00094 (test9.go:27)	MOVL	AX, "".d+28(SP)
	0x0062 00098 (test9.go:28)	MOVL	$3, 4(AX)
	0x0069 00105 (test9.go:28)	MOVL	runtime.writeBarrier(SB), CX
	0x006f 00111 (test9.go:28)	TESTL	CX, CX
	0x0071 00113 (test9.go:28)	JNE	$0, 568
	0x0077 00119 (test9.go:28)	LEAL	go.string."Tom"(SB), CX
	0x007d 00125 (test9.go:28)	MOVL	CX, (AX)
	0x007f 00127 (test9.go:29)	LEAL	go.itab.*"".Human,"".Speaker(SB), CX
	0x0085 00133 (test9.go:29)	MOVL	CX, (SP)
	0x0088 00136 (test9.go:29)	MOVL	"".h+24(SP), CX
	0x008c 00140 (test9.go:29)	MOVL	CX, 4(SP)
	0x0090 00144 (test9.go:29)	PCDATA	$0, $3
	0x0090 00144 (test9.go:29)	CALL	"".Say(SB)
	0x0095 00149 (test9.go:30)	LEAL	go.itab.*"".Dog,"".Speaker(SB), AX
	0x009b 00155 (test9.go:30)	MOVL	AX, (SP)
	0x009e 00158 (test9.go:30)	MOVL	"".d+28(SP), AX
	0x00a2 00162 (test9.go:30)	MOVL	AX, 4(SP)
	0x00a6 00166 (test9.go:30)	PCDATA	$0, $0
	0x00a6 00166 (test9.go:30)	CALL	"".Say(SB)
	0x00ab 00171 (test9.go:31)	LEAL	type.chan int(SB), AX
	0x00b1 00177 (test9.go:31)	MOVL	AX, (SP)
	0x00b4 00180 (test9.go:31)	MOVL	$0, 4(SP)
	0x00bc 00188 (test9.go:31)	MOVL	$0, 8(SP)
	0x00c4 00196 (test9.go:31)	PCDATA	$0, $0
	0x00c4 00196 (test9.go:31)	CALL	runtime.makechan(SB)
	0x00c9 00201 (test9.go:31)	MOVL	12(SP), AX
	0x00cd 00205 (test9.go:31)	MOVL	AX, "".ch+32(SP)
	0x00d1 00209 (test9.go:32)	MOVL	AX, 8(SP)
	0x00d5 00213 (test9.go:32)	MOVL	$4, (SP)
	0x00dc 00220 (test9.go:32)	LEAL	"".pump·f(SB), CX
	0x00e2 00226 (test9.go:32)	MOVL	CX, 4(SP)
	0x00e6 00230 (test9.go:32)	PCDATA	$0, $4
	0x00e6 00230 (test9.go:32)	CALL	runtime.newproc(SB)
	0x00eb 00235 (test9.go:33)	MOVL	"".ch+32(SP), AX
	0x00ef 00239 (test9.go:33)	MOVL	AX, 8(SP)
	0x00f3 00243 (test9.go:33)	MOVL	$4, (SP)
	0x00fa 00250 (test9.go:33)	LEAL	"".suck·f(SB), AX
	0x0100 00256 (test9.go:33)	MOVL	AX, 4(SP)
	0x0104 00260 (test9.go:33)	PCDATA	$0, $0
	0x0104 00260 (test9.go:33)	CALL	runtime.newproc(SB)
	0x0109 00265 (test9.go:34)	LEAL	type.int(SB), AX
	0x010f 00271 (test9.go:34)	MOVL	AX, (SP)
	0x0112 00274 (test9.go:34)	PCDATA	$0, $0
	0x0112 00274 (test9.go:34)	CALL	runtime.newobject(SB)
	0x0117 00279 (test9.go:34)	MOVL	4(SP), AX
	0x011b 00283 (test9.go:34)	MOVL	AX, ""..autotmp_17+36(SP)
	0x011f 00287 (test9.go:34)	MOVL	$4660, (AX)
	0x0125 00293 (test9.go:36)	MOVL	$0, ""..autotmp_3+72(SP)
	0x012d 00301 (test9.go:36)	MOVL	$0, ""..autotmp_3+76(SP)
	0x0135 00309 (test9.go:36)	LEAL	type.*[4]uint8(SB), CX
	0x013b 00315 (test9.go:36)	MOVL	CX, ""..autotmp_3+72(SP)
	0x013f 00319 (test9.go:36)	MOVL	AX, ""..autotmp_3+76(SP)
	0x0143 00323 (test9.go:36)	LEAL	""..autotmp_3+72(SP), CX
	0x0147 00327 (test9.go:36)	MOVL	CX, (SP)
	0x014a 00330 (test9.go:36)	MOVL	$1, 4(SP)
	0x0152 00338 (test9.go:36)	MOVL	$1, 8(SP)
	0x015a 00346 (test9.go:36)	PCDATA	$0, $5
	0x015a 00346 (test9.go:36)	CALL	fmt.Println(SB)
	0x015f 00351 (test9.go:38)	MOVL	""..autotmp_17+36(SP), AX
	0x0163 00355 (test9.go:38)	MOVBLZX	(AX), AX
	0x0166 00358 (test9.go:38)	TESTB	AX, AX
	0x0168 00360 (test9.go:38)	JNE	466
	0x016a 00362 (test9.go:39)	LEAL	go.string."本机器：大端"(SB), AX
	0x0170 00368 (test9.go:39)	MOVL	AX, ""..autotmp_6+56(SP)
	0x0174 00372 (test9.go:39)	MOVL	$18, ""..autotmp_6+60(SP)
	0x017c 00380 (test9.go:39)	MOVL	$0, ""..autotmp_5+64(SP)
	0x0184 00388 (test9.go:39)	MOVL	$0, ""..autotmp_5+68(SP)
	0x018c 00396 (test9.go:39)	LEAL	type.string(SB), AX
	0x0192 00402 (test9.go:39)	MOVL	AX, (SP)
	0x0195 00405 (test9.go:39)	LEAL	""..autotmp_6+56(SP), AX
	0x0199 00409 (test9.go:39)	MOVL	AX, 4(SP)
	0x019d 00413 (test9.go:39)	PCDATA	$0, $6
	0x019d 00413 (test9.go:39)	CALL	runtime.convT2E(SB)
	0x01a2 00418 (test9.go:39)	MOVL	8(SP), AX
	0x01a6 00422 (test9.go:39)	MOVL	12(SP), CX
	0x01aa 00426 (test9.go:39)	MOVL	AX, ""..autotmp_5+64(SP)
	0x01ae 00430 (test9.go:39)	MOVL	CX, ""..autotmp_5+68(SP)
	0x01b2 00434 (test9.go:39)	LEAL	""..autotmp_5+64(SP), AX
	0x01b6 00438 (test9.go:39)	MOVL	AX, (SP)
	0x01b9 00441 (test9.go:39)	MOVL	$1, 4(SP)
	0x01c1 00449 (test9.go:39)	MOVL	$1, 8(SP)
	0x01c9 00457 (test9.go:39)	PCDATA	$0, $6
	0x01c9 00457 (test9.go:39)	CALL	fmt.Println(SB)
	0x01ce 00462 (test9.go:44)	ADDL	$80, SP
	0x01d1 00465 (test9.go:44)	RET
	0x01d2 00466 (test9.go:41)	LEAL	go.string."本机器：小端"(SB), AX
	0x01d8 00472 (test9.go:41)	MOVL	AX, ""..autotmp_8+40(SP)
	0x01dc 00476 (test9.go:41)	MOVL	$18, ""..autotmp_8+44(SP)
	0x01e4 00484 (test9.go:41)	MOVL	$0, ""..autotmp_7+48(SP)
	0x01ec 00492 (test9.go:41)	MOVL	$0, ""..autotmp_7+52(SP)
	0x01f4 00500 (test9.go:39)	LEAL	type.string(SB), AX
	0x01fa 00506 (test9.go:41)	MOVL	AX, (SP)
	0x01fd 00509 (test9.go:41)	LEAL	""..autotmp_8+40(SP), AX
	0x0201 00513 (test9.go:41)	MOVL	AX, 4(SP)
	0x0205 00517 (test9.go:41)	PCDATA	$0, $7
	0x0205 00517 (test9.go:41)	CALL	runtime.convT2E(SB)
	0x020a 00522 (test9.go:41)	MOVL	12(SP), AX
	0x020e 00526 (test9.go:41)	MOVL	8(SP), CX
	0x0212 00530 (test9.go:41)	MOVL	CX, ""..autotmp_7+48(SP)
	0x0216 00534 (test9.go:41)	MOVL	AX, ""..autotmp_7+52(SP)
	0x021a 00538 (test9.go:41)	LEAL	""..autotmp_7+48(SP), AX
	0x021e 00542 (test9.go:41)	MOVL	AX, (SP)
	0x0221 00545 (test9.go:41)	MOVL	$1, 4(SP)
	0x0229 00553 (test9.go:41)	MOVL	$1, 8(SP)
	0x0231 00561 (test9.go:41)	PCDATA	$0, $7
	0x0231 00561 (test9.go:41)	CALL	fmt.Println(SB)
	0x0236 00566 (test9.go:44)	JMP	462
	0x0238 00568 (test9.go:28)	MOVL	AX, (SP)
	0x023b 00571 (test9.go:28)	LEAL	go.string."Tom"(SB), CX
	0x0241 00577 (test9.go:28)	MOVL	CX, 4(SP)
	0x0245 00581 (test9.go:28)	PCDATA	$0, $2
	0x0245 00581 (test9.go:28)	CALL	runtime.writebarrierptr(SB)
	0x024a 00586 (test9.go:30)	MOVL	"".d+28(SP), AX
	0x024e 00590 (test9.go:29)	JMP	127
	0x0253 00595 (test9.go:26)	MOVL	AX, (SP)
	0x0256 00598 (test9.go:26)	LEAL	go.string."Ben"(SB), CX
	0x025c 00604 (test9.go:26)	MOVL	CX, 4(SP)
	0x0260 00608 (test9.go:26)	PCDATA	$0, $1
	0x0260 00608 (test9.go:26)	CALL	runtime.writebarrierptr(SB)
	0x0265 00613 (test9.go:29)	MOVL	"".h+24(SP), AX
	0x0269 00617 (test9.go:27)	JMP	76
	0x026e 00622 (test9.go:27)	NOP
	0x026e 00622 (test9.go:24)	PCDATA	$0, $-1
	0x026e 00622 (test9.go:24)	CALL	runtime.morestack_noctxt(SB)
	0x0273 00627 (test9.go:24)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 0f 86 58 02 00 00 83 ec 50 8d 05 00 00 00 00 89  ..X.....P.......
	0x0020 04 24 e8 00 00 00 00 8b 44 24 04 89 44 24 18 c7  .$......D$..D$..
	0x0030 40 04 03 00 00 00 8b 0d 00 00 00 00 85 c9 0f 85  @...............
	0x0040 0f 02 00 00 8d 0d 00 00 00 00 89 08 8d 0d 00 00  ................
	0x0050 00 00 89 0c 24 e8 00 00 00 00 8b 44 24 04 89 44  ....$......D$..D
	0x0060 24 1c c7 40 04 03 00 00 00 8b 0d 00 00 00 00 85  $..@............
	0x0070 c9 0f 85 c1 01 00 00 8d 0d 00 00 00 00 89 08 8d  ................
	0x0080 0d 00 00 00 00 89 0c 24 8b 4c 24 18 89 4c 24 04  .......$.L$..L$.
	0x0090 e8 00 00 00 00 8d 05 00 00 00 00 89 04 24 8b 44  .............$.D
	0x00a0 24 1c 89 44 24 04 e8 00 00 00 00 8d 05 00 00 00  $..D$...........
	0x00b0 00 89 04 24 c7 44 24 04 00 00 00 00 c7 44 24 08  ...$.D$......D$.
	0x00c0 00 00 00 00 e8 00 00 00 00 8b 44 24 0c 89 44 24  ..........D$..D$
	0x00d0 20 89 44 24 08 c7 04 24 04 00 00 00 8d 0d 00 00   .D$...$........
	0x00e0 00 00 89 4c 24 04 e8 00 00 00 00 8b 44 24 20 89  ...L$.......D$ .
	0x00f0 44 24 08 c7 04 24 04 00 00 00 8d 05 00 00 00 00  D$...$..........
	0x0100 89 44 24 04 e8 00 00 00 00 8d 05 00 00 00 00 89  .D$.............
	0x0110 04 24 e8 00 00 00 00 8b 44 24 04 89 44 24 24 c7  .$......D$..D$$.
	0x0120 00 34 12 00 00 c7 44 24 48 00 00 00 00 c7 44 24  .4....D$H.....D$
	0x0130 4c 00 00 00 00 8d 0d 00 00 00 00 89 4c 24 48 89  L...........L$H.
	0x0140 44 24 4c 8d 4c 24 48 89 0c 24 c7 44 24 04 01 00  D$L.L$H..$.D$...
	0x0150 00 00 c7 44 24 08 01 00 00 00 e8 00 00 00 00 8b  ...D$...........
	0x0160 44 24 24 0f b6 00 84 c0 75 68 8d 05 00 00 00 00  D$$.....uh......
	0x0170 89 44 24 38 c7 44 24 3c 12 00 00 00 c7 44 24 40  .D$8.D$<.....D$@
	0x0180 00 00 00 00 c7 44 24 44 00 00 00 00 8d 05 00 00  .....D$D........
	0x0190 00 00 89 04 24 8d 44 24 38 89 44 24 04 e8 00 00  ....$.D$8.D$....
	0x01a0 00 00 8b 44 24 08 8b 4c 24 0c 89 44 24 40 89 4c  ...D$..L$..D$@.L
	0x01b0 24 44 8d 44 24 40 89 04 24 c7 44 24 04 01 00 00  $D.D$@..$.D$....
	0x01c0 00 c7 44 24 08 01 00 00 00 e8 00 00 00 00 83 c4  ..D$............
	0x01d0 50 c3 8d 05 00 00 00 00 89 44 24 28 c7 44 24 2c  P........D$(.D$,
	0x01e0 12 00 00 00 c7 44 24 30 00 00 00 00 c7 44 24 34  .....D$0.....D$4
	0x01f0 00 00 00 00 8d 05 00 00 00 00 89 04 24 8d 44 24  ............$.D$
	0x0200 28 89 44 24 04 e8 00 00 00 00 8b 44 24 0c 8b 4c  (.D$.......D$..L
	0x0210 24 08 89 4c 24 30 89 44 24 34 8d 44 24 30 89 04  $..L$0.D$4.D$0..
	0x0220 24 c7 44 24 04 01 00 00 00 c7 44 24 08 01 00 00  $.D$......D$....
	0x0230 00 e8 00 00 00 00 eb 96 89 04 24 8d 0d 00 00 00  ..........$.....
	0x0240 00 89 4c 24 04 e8 00 00 00 00 8b 44 24 1c e9 2c  ..L$.......D$..,
	0x0250 fe ff ff 89 04 24 8d 0d 00 00 00 00 89 4c 24 04  .....$.......L$.
	0x0260 e8 00 00 00 00 8b 44 24 18 e9 de fd ff ff e8 00  ......D$........
	0x0270 00 00 00 e9 88 fd ff ff                          ........
	rel 9+4 t=16 TLS+0
	rel 27+4 t=1 type."".Human+0
	rel 35+4 t=8 runtime.newobject+0
	rel 56+4 t=1 runtime.writeBarrier+0
	rel 70+4 t=1 go.string."Ben"+0
	rel 78+4 t=1 type."".Dog+0
	rel 86+4 t=8 runtime.newobject+0
	rel 107+4 t=1 runtime.writeBarrier+0
	rel 121+4 t=1 go.string."Tom"+0
	rel 129+4 t=1 go.itab.*"".Human,"".Speaker+0
	rel 145+4 t=8 "".Say+0
	rel 151+4 t=1 go.itab.*"".Dog,"".Speaker+0
	rel 167+4 t=8 "".Say+0
	rel 173+4 t=1 type.chan int+0
	rel 197+4 t=8 runtime.makechan+0
	rel 222+4 t=1 "".pump·f+0
	rel 231+4 t=8 runtime.newproc+0
	rel 252+4 t=1 "".suck·f+0
	rel 261+4 t=8 runtime.newproc+0
	rel 267+4 t=1 type.int+0
	rel 275+4 t=8 runtime.newobject+0
	rel 311+4 t=1 type.*[4]uint8+0
	rel 347+4 t=8 fmt.Println+0
	rel 364+4 t=1 go.string."本机器：大端"+0
	rel 398+4 t=1 type.string+0
	rel 414+4 t=8 runtime.convT2E+0
	rel 458+4 t=8 fmt.Println+0
	rel 468+4 t=1 go.string."本机器：小端"+0
	rel 502+4 t=1 type.string+0
	rel 518+4 t=8 runtime.convT2E+0
	rel 562+4 t=8 fmt.Println+0
	rel 573+4 t=1 go.string."Tom"+0
	rel 582+4 t=8 runtime.writebarrierptr+0
	rel 600+4 t=1 go.string."Ben"+0
	rel 609+4 t=8 runtime.writebarrierptr+0
	rel 623+4 t=8 runtime.morestack_noctxt+0
"".Say t=1 size=48 args=0x8 locals=0x4
	0x0000 00000 (test9.go:46)	TEXT	"".Say(SB), $4-8
	0x0000 00000 (test9.go:46)	MOVL	TLS, CX
	0x0007 00007 (test9.go:46)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:46)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:46)	JLS	41
	0x0012 00018 (test9.go:46)	SUBL	$4, SP
	0x0015 00021 (test9.go:46)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0015 00021 (test9.go:46)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0015 00021 (test9.go:47)	MOVL	"".s+8(FP), AX
	0x0019 00025 (test9.go:47)	MOVL	20(AX), AX
	0x001c 00028 (test9.go:47)	MOVL	"".s+12(FP), CX
	0x0020 00032 (test9.go:47)	MOVL	CX, (SP)
	0x0023 00035 (test9.go:47)	PCDATA	$0, $1
	0x0023 00035 (test9.go:47)	CALL	AX
	0x0025 00037 (test9.go:48)	ADDL	$4, SP
	0x0028 00040 (test9.go:48)	RET
	0x0029 00041 (test9.go:48)	NOP
	0x0029 00041 (test9.go:46)	PCDATA	$0, $-1
	0x0029 00041 (test9.go:46)	CALL	runtime.morestack_noctxt(SB)
	0x002e 00046 (test9.go:46)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 76 17 83 ec 04 8b 44 24 08 8b 40 14 8b 4c 24 0c  v.....D$..@..L$.
	0x0020 89 0c 24 ff d0 83 c4 04 c3 e8 00 00 00 00 eb d0  ..$.............
	rel 9+4 t=16 TLS+0
	rel 35+0 t=11 +0
	rel 42+4 t=8 runtime.morestack_noctxt+0
"".(*Human).Say t=1 size=184 args=0x4 locals=0x2c
	0x0000 00000 (test9.go:50)	TEXT	"".(*Human).Say(SB), $44-4
	0x0000 00000 (test9.go:50)	MOVL	TLS, CX
	0x0007 00007 (test9.go:50)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:50)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:50)	JLS	174
	0x0016 00022 (test9.go:50)	SUBL	$44, SP
	0x0019 00025 (test9.go:50)	FUNCDATA	$0, gclocals·2d7c1615616d4cf40d01b3385155ed6e(SB)
	0x0019 00025 (test9.go:50)	FUNCDATA	$1, gclocals·488112b86ce2e5d099bbd25cfb5a88fa(SB)
	0x0019 00025 (test9.go:51)	MOVL	"".h+48(FP), AX
	0x001d 00029 (test9.go:51)	MOVL	(AX), CX
	0x001f 00031 (test9.go:51)	MOVL	4(AX), AX
	0x0022 00034 (test9.go:51)	MOVL	CX, 4(SP)
	0x0026 00038 (test9.go:51)	MOVL	AX, 8(SP)
	0x002a 00042 (test9.go:51)	MOVL	$0, (SP)
	0x0031 00049 (test9.go:51)	LEAL	go.string." Say Hello Go"(SB), AX
	0x0037 00055 (test9.go:51)	MOVL	AX, 12(SP)
	0x003b 00059 (test9.go:51)	MOVL	$13, 16(SP)
	0x0043 00067 (test9.go:51)	PCDATA	$0, $1
	0x0043 00067 (test9.go:51)	CALL	runtime.concatstring2(SB)
	0x0048 00072 (test9.go:51)	MOVL	24(SP), AX
	0x004c 00076 (test9.go:51)	MOVL	20(SP), CX
	0x0050 00080 (test9.go:51)	MOVL	CX, ""..autotmp_19+28(SP)
	0x0054 00084 (test9.go:51)	MOVL	AX, ""..autotmp_19+32(SP)
	0x0058 00088 (test9.go:51)	MOVL	$0, ""..autotmp_18+36(SP)
	0x0060 00096 (test9.go:51)	MOVL	$0, ""..autotmp_18+40(SP)
	0x0068 00104 (test9.go:51)	LEAL	type.string(SB), AX
	0x006e 00110 (test9.go:51)	MOVL	AX, (SP)
	0x0071 00113 (test9.go:51)	LEAL	""..autotmp_19+28(SP), AX
	0x0075 00117 (test9.go:51)	MOVL	AX, 4(SP)
	0x0079 00121 (test9.go:51)	PCDATA	$0, $2
	0x0079 00121 (test9.go:51)	CALL	runtime.convT2E(SB)
	0x007e 00126 (test9.go:51)	MOVL	8(SP), AX
	0x0082 00130 (test9.go:51)	MOVL	12(SP), CX
	0x0086 00134 (test9.go:51)	MOVL	AX, ""..autotmp_18+36(SP)
	0x008a 00138 (test9.go:51)	MOVL	CX, ""..autotmp_18+40(SP)
	0x008e 00142 (test9.go:51)	LEAL	""..autotmp_18+36(SP), AX
	0x0092 00146 (test9.go:51)	MOVL	AX, (SP)
	0x0095 00149 (test9.go:51)	MOVL	$1, 4(SP)
	0x009d 00157 (test9.go:51)	MOVL	$1, 8(SP)
	0x00a5 00165 (test9.go:51)	PCDATA	$0, $2
	0x00a5 00165 (test9.go:51)	CALL	fmt.Println(SB)
	0x00aa 00170 (test9.go:52)	ADDL	$44, SP
	0x00ad 00173 (test9.go:52)	RET
	0x00ae 00174 (test9.go:52)	NOP
	0x00ae 00174 (test9.go:50)	PCDATA	$0, $-1
	0x00ae 00174 (test9.go:50)	CALL	runtime.morestack_noctxt(SB)
	0x00b3 00179 (test9.go:50)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 0f 86 98 00 00 00 83 ec 2c 8b 44 24 30 8b 08 8b  ........,.D$0...
	0x0020 40 04 89 4c 24 04 89 44 24 08 c7 04 24 00 00 00  @..L$..D$...$...
	0x0030 00 8d 05 00 00 00 00 89 44 24 0c c7 44 24 10 0d  ........D$..D$..
	0x0040 00 00 00 e8 00 00 00 00 8b 44 24 18 8b 4c 24 14  .........D$..L$.
	0x0050 89 4c 24 1c 89 44 24 20 c7 44 24 24 00 00 00 00  .L$..D$ .D$$....
	0x0060 c7 44 24 28 00 00 00 00 8d 05 00 00 00 00 89 04  .D$(............
	0x0070 24 8d 44 24 1c 89 44 24 04 e8 00 00 00 00 8b 44  $.D$..D$.......D
	0x0080 24 08 8b 4c 24 0c 89 44 24 24 89 4c 24 28 8d 44  $..L$..D$$.L$(.D
	0x0090 24 24 89 04 24 c7 44 24 04 01 00 00 00 c7 44 24  $$..$.D$......D$
	0x00a0 08 01 00 00 00 e8 00 00 00 00 83 c4 2c c3 e8 00  ............,...
	0x00b0 00 00 00 e9 48 ff ff ff                          ....H...
	rel 9+4 t=16 TLS+0
	rel 51+4 t=1 go.string." Say Hello Go"+0
	rel 68+4 t=8 runtime.concatstring2+0
	rel 106+4 t=1 type.string+0
	rel 122+4 t=8 runtime.convT2E+0
	rel 166+4 t=8 fmt.Println+0
	rel 175+4 t=8 runtime.morestack_noctxt+0
"".(*Dog).Say t=1 size=184 args=0x4 locals=0x2c
	0x0000 00000 (test9.go:54)	TEXT	"".(*Dog).Say(SB), $44-4
	0x0000 00000 (test9.go:54)	MOVL	TLS, CX
	0x0007 00007 (test9.go:54)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:54)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:54)	JLS	174
	0x0016 00022 (test9.go:54)	SUBL	$44, SP
	0x0019 00025 (test9.go:54)	FUNCDATA	$0, gclocals·2d7c1615616d4cf40d01b3385155ed6e(SB)
	0x0019 00025 (test9.go:54)	FUNCDATA	$1, gclocals·488112b86ce2e5d099bbd25cfb5a88fa(SB)
	0x0019 00025 (test9.go:55)	MOVL	"".d+48(FP), AX
	0x001d 00029 (test9.go:55)	MOVL	(AX), CX
	0x001f 00031 (test9.go:55)	MOVL	4(AX), AX
	0x0022 00034 (test9.go:55)	MOVL	CX, 4(SP)
	0x0026 00038 (test9.go:55)	MOVL	AX, 8(SP)
	0x002a 00042 (test9.go:55)	MOVL	$0, (SP)
	0x0031 00049 (test9.go:55)	LEAL	go.string." Say Hello Go"(SB), AX
	0x0037 00055 (test9.go:55)	MOVL	AX, 12(SP)
	0x003b 00059 (test9.go:55)	MOVL	$13, 16(SP)
	0x0043 00067 (test9.go:55)	PCDATA	$0, $1
	0x0043 00067 (test9.go:55)	CALL	runtime.concatstring2(SB)
	0x0048 00072 (test9.go:55)	MOVL	24(SP), AX
	0x004c 00076 (test9.go:55)	MOVL	20(SP), CX
	0x0050 00080 (test9.go:55)	MOVL	CX, ""..autotmp_24+28(SP)
	0x0054 00084 (test9.go:55)	MOVL	AX, ""..autotmp_24+32(SP)
	0x0058 00088 (test9.go:55)	MOVL	$0, ""..autotmp_23+36(SP)
	0x0060 00096 (test9.go:55)	MOVL	$0, ""..autotmp_23+40(SP)
	0x0068 00104 (test9.go:55)	LEAL	type.string(SB), AX
	0x006e 00110 (test9.go:55)	MOVL	AX, (SP)
	0x0071 00113 (test9.go:55)	LEAL	""..autotmp_24+28(SP), AX
	0x0075 00117 (test9.go:55)	MOVL	AX, 4(SP)
	0x0079 00121 (test9.go:55)	PCDATA	$0, $2
	0x0079 00121 (test9.go:55)	CALL	runtime.convT2E(SB)
	0x007e 00126 (test9.go:55)	MOVL	8(SP), AX
	0x0082 00130 (test9.go:55)	MOVL	12(SP), CX
	0x0086 00134 (test9.go:55)	MOVL	AX, ""..autotmp_23+36(SP)
	0x008a 00138 (test9.go:55)	MOVL	CX, ""..autotmp_23+40(SP)
	0x008e 00142 (test9.go:55)	LEAL	""..autotmp_23+36(SP), AX
	0x0092 00146 (test9.go:55)	MOVL	AX, (SP)
	0x0095 00149 (test9.go:55)	MOVL	$1, 4(SP)
	0x009d 00157 (test9.go:55)	MOVL	$1, 8(SP)
	0x00a5 00165 (test9.go:55)	PCDATA	$0, $2
	0x00a5 00165 (test9.go:55)	CALL	fmt.Println(SB)
	0x00aa 00170 (test9.go:56)	ADDL	$44, SP
	0x00ad 00173 (test9.go:56)	RET
	0x00ae 00174 (test9.go:56)	NOP
	0x00ae 00174 (test9.go:54)	PCDATA	$0, $-1
	0x00ae 00174 (test9.go:54)	CALL	runtime.morestack_noctxt(SB)
	0x00b3 00179 (test9.go:54)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 0f 86 98 00 00 00 83 ec 2c 8b 44 24 30 8b 08 8b  ........,.D$0...
	0x0020 40 04 89 4c 24 04 89 44 24 08 c7 04 24 00 00 00  @..L$..D$...$...
	0x0030 00 8d 05 00 00 00 00 89 44 24 0c c7 44 24 10 0d  ........D$..D$..
	0x0040 00 00 00 e8 00 00 00 00 8b 44 24 18 8b 4c 24 14  .........D$..L$.
	0x0050 89 4c 24 1c 89 44 24 20 c7 44 24 24 00 00 00 00  .L$..D$ .D$$....
	0x0060 c7 44 24 28 00 00 00 00 8d 05 00 00 00 00 89 04  .D$(............
	0x0070 24 8d 44 24 1c 89 44 24 04 e8 00 00 00 00 8b 44  $.D$..D$.......D
	0x0080 24 08 8b 4c 24 0c 89 44 24 24 89 4c 24 28 8d 44  $..L$..D$$.L$(.D
	0x0090 24 24 89 04 24 c7 44 24 04 01 00 00 00 c7 44 24  $$..$.D$......D$
	0x00a0 08 01 00 00 00 e8 00 00 00 00 83 c4 2c c3 e8 00  ............,...
	0x00b0 00 00 00 e9 48 ff ff ff                          ....H...
	rel 9+4 t=16 TLS+0
	rel 51+4 t=1 go.string." Say Hello Go"+0
	rel 68+4 t=8 runtime.concatstring2+0
	rel 106+4 t=1 type.string+0
	rel 122+4 t=8 runtime.convT2E+0
	rel 166+4 t=8 fmt.Println+0
	rel 175+4 t=8 runtime.morestack_noctxt+0
"".pump t=1 size=75 args=0x4 locals=0x14
	0x0000 00000 (test9.go:57)	TEXT	"".pump(SB), $20-4
	0x0000 00000 (test9.go:57)	MOVL	TLS, CX
	0x0007 00007 (test9.go:57)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:57)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:57)	JLS	68
	0x0012 00018 (test9.go:57)	SUBL	$20, SP
	0x0015 00021 (test9.go:57)	FUNCDATA	$0, gclocals·a36216b97439c93dafebe03e7f0808b5(SB)
	0x0015 00021 (test9.go:57)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0015 00021 (test9.go:58)	MOVL	$0, AX
	0x0017 00023 (test9.go:59)	MOVL	AX, "".i+12(SP)
	0x001b 00027 (test9.go:59)	MOVL	AX, ""..autotmp_28+16(SP)
	0x001f 00031 (test9.go:59)	LEAL	type.chan int(SB), CX
	0x0025 00037 (test9.go:59)	MOVL	CX, (SP)
	0x0028 00040 (test9.go:59)	MOVL	"".ch+24(FP), DX
	0x002c 00044 (test9.go:59)	MOVL	DX, 4(SP)
	0x0030 00048 (test9.go:59)	LEAL	""..autotmp_28+16(SP), BX
	0x0034 00052 (test9.go:59)	MOVL	BX, 8(SP)
	0x0038 00056 (test9.go:59)	PCDATA	$0, $0
	0x0038 00056 (test9.go:59)	CALL	runtime.chansend1(SB)
	0x003d 00061 (test9.go:58)	MOVL	"".i+12(SP), AX
	0x0041 00065 (test9.go:58)	INCL	AX
	0x0042 00066 (test9.go:59)	JMP	23
	0x0044 00068 (test9.go:59)	NOP
	0x0044 00068 (test9.go:57)	PCDATA	$0, $-1
	0x0044 00068 (test9.go:57)	CALL	runtime.morestack_noctxt(SB)
	0x0049 00073 (test9.go:57)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 76 32 83 ec 14 31 c0 89 44 24 0c 89 44 24 10 8d  v2...1..D$..D$..
	0x0020 0d 00 00 00 00 89 0c 24 8b 54 24 18 89 54 24 04  .......$.T$..T$.
	0x0030 8d 5c 24 10 89 5c 24 08 e8 00 00 00 00 8b 44 24  .\$..\$.......D$
	0x0040 0c 40 eb d3 e8 00 00 00 00 eb b5                 .@.........
	rel 9+4 t=16 TLS+0
	rel 33+4 t=1 type.chan int+0
	rel 57+4 t=8 runtime.chansend1+0
	rel 69+4 t=8 runtime.morestack_noctxt+0
"".suck t=1 size=153 args=0x4 locals=0x24
	0x0000 00000 (test9.go:63)	TEXT	"".suck(SB), $36-4
	0x0000 00000 (test9.go:63)	MOVL	TLS, CX
	0x0007 00007 (test9.go:63)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:63)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:63)	JLS	143
	0x0012 00018 (test9.go:63)	SUBL	$36, SP
	0x0015 00021 (test9.go:63)	FUNCDATA	$0, gclocals·09b80ec389a9e6ac09cfa1cd3c45263d(SB)
	0x0015 00021 (test9.go:63)	FUNCDATA	$1, gclocals·e226d4ae4a7cad8835311c6a4683c14f(SB)
	0x0015 00021 (test9.go:65)	MOVL	$0, ""..autotmp_31+24(SP)
	0x001d 00029 (test9.go:65)	LEAL	type.chan int(SB), AX
	0x0023 00035 (test9.go:65)	MOVL	AX, (SP)
	0x0026 00038 (test9.go:65)	MOVL	"".ch+40(FP), CX
	0x002a 00042 (test9.go:65)	MOVL	CX, 4(SP)
	0x002e 00046 (test9.go:65)	LEAL	""..autotmp_31+24(SP), DX
	0x0032 00050 (test9.go:65)	MOVL	DX, 8(SP)
	0x0036 00054 (test9.go:65)	PCDATA	$0, $0
	0x0036 00054 (test9.go:65)	CALL	runtime.chanrecv1(SB)
	0x003b 00059 (test9.go:65)	MOVL	$0, ""..autotmp_30+28(SP)
	0x0043 00067 (test9.go:65)	MOVL	$0, ""..autotmp_30+32(SP)
	0x004b 00075 (test9.go:65)	LEAL	type.int(SB), AX
	0x0051 00081 (test9.go:65)	MOVL	AX, (SP)
	0x0054 00084 (test9.go:65)	LEAL	""..autotmp_31+24(SP), CX
	0x0058 00088 (test9.go:65)	MOVL	CX, 4(SP)
	0x005c 00092 (test9.go:65)	PCDATA	$0, $1
	0x005c 00092 (test9.go:65)	CALL	runtime.convT2E(SB)
	0x0061 00097 (test9.go:65)	MOVL	8(SP), AX
	0x0065 00101 (test9.go:65)	MOVL	12(SP), CX
	0x0069 00105 (test9.go:65)	MOVL	AX, ""..autotmp_30+28(SP)
	0x006d 00109 (test9.go:65)	MOVL	CX, ""..autotmp_30+32(SP)
	0x0071 00113 (test9.go:65)	LEAL	""..autotmp_30+28(SP), AX
	0x0075 00117 (test9.go:65)	MOVL	AX, (SP)
	0x0078 00120 (test9.go:65)	MOVL	$1, 4(SP)
	0x0080 00128 (test9.go:65)	MOVL	$1, 8(SP)
	0x0088 00136 (test9.go:65)	PCDATA	$0, $1
	0x0088 00136 (test9.go:65)	CALL	fmt.Println(SB)
	0x008d 00141 (test9.go:65)	JMP	21
	0x008f 00143 (test9.go:65)	NOP
	0x008f 00143 (test9.go:63)	PCDATA	$0, $-1
	0x008f 00143 (test9.go:63)	CALL	runtime.morestack_noctxt(SB)
	0x0094 00148 (test9.go:63)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 76 7d 83 ec 24 c7 44 24 18 00 00 00 00 8d 05 00  v}..$.D$........
	0x0020 00 00 00 89 04 24 8b 4c 24 28 89 4c 24 04 8d 54  .....$.L$(.L$..T
	0x0030 24 18 89 54 24 08 e8 00 00 00 00 c7 44 24 1c 00  $..T$.......D$..
	0x0040 00 00 00 c7 44 24 20 00 00 00 00 8d 05 00 00 00  ....D$ .........
	0x0050 00 89 04 24 8d 4c 24 18 89 4c 24 04 e8 00 00 00  ...$.L$..L$.....
	0x0060 00 8b 44 24 08 8b 4c 24 0c 89 44 24 1c 89 4c 24  ..D$..L$..D$..L$
	0x0070 20 8d 44 24 1c 89 04 24 c7 44 24 04 01 00 00 00   .D$...$.D$.....
	0x0080 c7 44 24 08 01 00 00 00 e8 00 00 00 00 eb 86 e8  .D$.............
	0x0090 00 00 00 00 e9 67 ff ff ff                       .....g...
	rel 9+4 t=16 TLS+0
	rel 31+4 t=1 type.chan int+0
	rel 55+4 t=8 runtime.chanrecv1+0
	rel 77+4 t=1 type.int+0
	rel 93+4 t=8 runtime.convT2E+0
	rel 137+4 t=8 fmt.Println+0
	rel 144+4 t=8 runtime.morestack_noctxt+0
"".init t=1 size=72 args=0x0 locals=0x0
	0x0000 00000 (test9.go:68)	TEXT	"".init(SB), $0-0
	0x0000 00000 (test9.go:68)	MOVL	TLS, CX
	0x0007 00007 (test9.go:68)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test9.go:68)	CMPL	SP, 8(CX)
	0x0010 00016 (test9.go:68)	JLS	65
	0x0012 00018 (test9.go:68)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (test9.go:68)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (test9.go:68)	MOVBLZX	"".initdone·(SB), AX
	0x0019 00025 (test9.go:68)	CMPB	AX, $1
	0x001c 00028 (test9.go:68)	JLS	$0, 31
	0x001e 00030 (test9.go:68)	RET
	0x001f 00031 (test9.go:68)	JNE	$0, 40
	0x0021 00033 (test9.go:68)	PCDATA	$0, $0
	0x0021 00033 (test9.go:68)	CALL	runtime.throwinit(SB)
	0x0026 00038 (test9.go:68)	UNDEF
	0x0028 00040 (test9.go:68)	MOVB	$1, "".initdone·(SB)
	0x002f 00047 (test9.go:68)	PCDATA	$0, $0
	0x002f 00047 (test9.go:68)	CALL	fmt.init(SB)
	0x0034 00052 (test9.go:68)	PCDATA	$0, $0
	0x0034 00052 (test9.go:68)	CALL	"".init.1(SB)
	0x0039 00057 (test9.go:68)	MOVB	$2, "".initdone·(SB)
	0x0040 00064 (test9.go:68)	RET
	0x0041 00065 (test9.go:68)	NOP
	0x0041 00065 (test9.go:68)	PCDATA	$0, $-1
	0x0041 00065 (test9.go:68)	CALL	runtime.morestack_noctxt(SB)
	0x0046 00070 (test9.go:68)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 76 2f 0f b6 05 00 00 00 00 80 f8 01 76 01 c3 75  v/..........v..u
	0x0020 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01 e8  ................
	0x0030 00 00 00 00 e8 00 00 00 00 c6 05 00 00 00 00 02  ................
	0x0040 c3 e8 00 00 00 00 eb b8                          ........
	rel 9+4 t=16 TLS+0
	rel 21+4 t=1 "".initdone·+0
	rel 34+4 t=8 runtime.throwinit+0
	rel 42+4 t=1 "".initdone·+0
	rel 48+4 t=8 fmt.init+0
	rel 53+4 t=8 "".init.1+0
	rel 59+4 t=1 "".initdone·+0
	rel 66+4 t=8 runtime.morestack_noctxt+0
"".Speaker.Say t=1 dupok size=65 args=0x8 locals=0x4
	0x0000 00000 (<autogenerated>:1)	TEXT	"".Speaker.Say(SB), $4-8
	0x0000 00000 (<autogenerated>:1)	MOVL	TLS, CX
	0x0007 00007 (<autogenerated>:1)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (<autogenerated>:1)	CMPL	SP, 8(CX)
	0x0010 00016 (<autogenerated>:1)	JLS	58
	0x0012 00018 (<autogenerated>:1)	SUBL	$4, SP
	0x0015 00021 (<autogenerated>:1)	MOVL	16(CX), BX
	0x0018 00024 (<autogenerated>:1)	TESTL	BX, BX
	0x001a 00026 (<autogenerated>:1)	JEQ	38
	0x001c 00028 (<autogenerated>:1)	LEAL	8(SP), DI
	0x0020 00032 (<autogenerated>:1)	CMPL	(BX), DI
	0x0022 00034 (<autogenerated>:1)	JNE	38
	0x0024 00036 (<autogenerated>:1)	MOVL	SP, (BX)
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (<autogenerated>:1)	MOVL	""..this+8(FP), AX
	0x002a 00042 (<autogenerated>:1)	MOVL	""..this+12(FP), CX
	0x002e 00046 (<autogenerated>:1)	MOVL	20(AX), AX
	0x0031 00049 (<autogenerated>:1)	MOVL	CX, (SP)
	0x0034 00052 (<autogenerated>:1)	PCDATA	$0, $1
	0x0034 00052 (<autogenerated>:1)	CALL	AX
	0x0036 00054 (<autogenerated>:1)	ADDL	$4, SP
	0x0039 00057 (<autogenerated>:1)	RET
	0x003a 00058 (<autogenerated>:1)	NOP
	0x003a 00058 (<autogenerated>:1)	PCDATA	$0, $-1
	0x003a 00058 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x003f 00063 (<autogenerated>:1)	JMP	0
	0x0000 64 8b 0d 14 00 00 00 8b 89 00 00 00 00 3b 61 08  d............;a.
	0x0010 76 28 83 ec 04 8b 59 10 85 db 74 0a 8d 7c 24 08  v(....Y...t..|$.
	0x0020 39 3b 75 02 89 23 8b 44 24 08 8b 4c 24 0c 8b 40  9;u..#.D$..L$..@
	0x0030 14 89 0c 24 ff d0 83 c4 04 c3 e8 00 00 00 00 eb  ...$............
	0x0040 bf                                               .
	rel 9+4 t=16 TLS+0
	rel 52+0 t=11 +0
	rel 59+4 t=8 runtime.morestack_noctxt+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb t=8 dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
go.info."".init.1 t=45 size=21
	0x0000 02 22 22 2e 69 6e 69 74 2e 31 00 00 00 00 00 00  ."".init.1......
	0x0010 00 00 00 01 00                                   .....
	rel 11+4 t=1 "".init.1+0
	rel 15+4 t=1 "".init.1+26
go.string."Ben" t=8 dupok size=3
	0x0000 42 65 6e                                         Ben
go.string."Tom" t=8 dupok size=3
	0x0000 54 6f 6d                                         Tom
go.string."本机器：大端" t=8 dupok size=18
	0x0000 e6 9c ac e6 9c ba e5 99 a8 ef bc 9a e5 a4 a7 e7  ................
	0x0010 ab af                                            ..
go.string."本机器：小端" t=8 dupok size=18
	0x0000 e6 9c ac e6 9c ba e5 99 a8 ef bc 9a e5 b0 8f e7  ................
	0x0010 ab af                                            ..
gclocals·885e690be6afa9a2a3e0e4279d6dbcbc t=8 dupok size=24
	0x0000 08 00 00 00 0e 00 00 00 00 00 01 00 03 00 02 00  ................
	0x0010 04 00 08 30 00 0d d0 00                          ...0....
gclocals·0ce64bbc7cfa5ef04d41c861de81a3d7 t=8 dupok size=8
	0x0000 08 00 00 00 00 00 00 00                          ........
go.info."".main t=45 size=56
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 01 04 68 00 04 9c 11 44 22 00 00 00 00 04 64  ...h....D".....d
	0x0020 00 04 9c 11 48 22 00 00 00 00 04 63 68 00 04 9c  ....H".....ch...
	0x0030 11 4c 22 00 00 00 00 00                          .L".....
	rel 9+4 t=1 "".main+0
	rel 13+4 t=1 "".main+632
	rel 26+4 t=28 go.info.*"".Human+0
	rel 38+4 t=28 go.info.*"".Dog+0
	rel 51+4 t=28 go.info.chan int+0
gclocals·69c1753bd5f81501d95132d08af04464 t=8 dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·dc9b0298814590ca3ffc3a889546fc8b t=8 dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
go.info."".Say t=45 size=27
	0x0000 02 22 22 2e 53 61 79 00 00 00 00 00 00 00 00 00  ."".Say.........
	0x0010 01 05 73 00 01 9c 00 00 00 00 00                 ..s........
	rel 8+4 t=1 "".Say+0
	rel 12+4 t=1 "".Say+48
	rel 22+4 t=28 go.info."".Speaker+0
go.string." Say Hello Go" t=8 dupok size=13
	0x0000 20 53 61 79 20 48 65 6c 6c 6f 20 47 6f            Say Hello Go
gclocals·488112b86ce2e5d099bbd25cfb5a88fa t=8 dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 00 0d                 ...........
gclocals·2d7c1615616d4cf40d01b3385155ed6e t=8 dupok size=11
	0x0000 03 00 00 00 01 00 00 00 01 00 00                 ...........
go.info."".(*Human).Say t=45 size=36
	0x0000 02 22 22 2e 28 2a 48 75 6d 61 6e 29 2e 53 61 79  ."".(*Human).Say
	0x0010 00 00 00 00 00 00 00 00 00 01 05 68 00 01 9c 00  ...........h....
	0x0020 00 00 00 00                                      ....
	rel 17+4 t=1 "".(*Human).Say+0
	rel 21+4 t=1 "".(*Human).Say+184
	rel 31+4 t=28 go.info.*"".Human+0
go.info."".(*Dog).Say t=45 size=34
	0x0000 02 22 22 2e 28 2a 44 6f 67 29 2e 53 61 79 00 00  ."".(*Dog).Say..
	0x0010 00 00 00 00 00 00 00 01 05 64 00 01 9c 00 00 00  .........d......
	0x0020 00 00                                            ..
	rel 15+4 t=1 "".(*Dog).Say+0
	rel 19+4 t=1 "".(*Dog).Say+184
	rel 29+4 t=28 go.info.*"".Dog+0
gclocals·a36216b97439c93dafebe03e7f0808b5 t=8 dupok size=9
	0x0000 01 00 00 00 01 00 00 00 01                       .........
go.info."".pump t=45 size=41
	0x0000 02 22 22 2e 70 75 6d 70 00 00 00 00 00 00 00 00  ."".pump........
	0x0010 00 01 04 69 00 04 9c 11 74 22 00 00 00 00 05 63  ...i....t".....c
	0x0020 68 00 01 9c 00 00 00 00 00                       h........
	rel 9+4 t=1 "".pump+0
	rel 13+4 t=1 "".pump+75
	rel 26+4 t=28 go.info.int+0
	rel 36+4 t=28 go.info.chan int+0
gclocals·e226d4ae4a7cad8835311c6a4683c14f t=8 dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 03                    ..........
gclocals·09b80ec389a9e6ac09cfa1cd3c45263d t=8 dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 01                    ..........
go.info."".suck t=45 size=29
	0x0000 02 22 22 2e 73 75 63 6b 00 00 00 00 00 00 00 00  ."".suck........
	0x0010 00 01 05 63 68 00 01 9c 00 00 00 00 00           ...ch........
	rel 9+4 t=1 "".suck+0
	rel 13+4 t=1 "".suck+153
	rel 24+4 t=28 go.info.chan int+0
go.info."".init t=45 size=19
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 01 00                                         ...
	rel 9+4 t=1 "".init+0
	rel 13+4 t=1 "".init+72
"".initdone· t=32 size=1
"".pump·f t=8 dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 "".pump+0
"".suck·f t=8 dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 "".suck+0
runtime.gcbits.01 t=8 dupok size=1
	0x0000 01                                               .
type..namedata.**main.Human. t=8 dupok size=15
	0x0000 00 00 0c 2a 2a 6d 61 69 6e 2e 48 75 6d 61 6e     ...**main.Human
type.**"".Human t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 92 a1 a9 be 00 04 04 36  ...............6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.**main.Human.+0
	rel 32+4 t=1 type.*"".Human+0
type..namedata.*main.Human. t=8 dupok size=14
	0x0000 01 00 0b 2a 6d 61 69 6e 2e 48 75 6d 61 6e        ...*main.Human
type..namedata.*func(*main.Human). t=8 dupok size=21
	0x0000 00 00 12 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 48  ...*func(*main.H
	0x0010 75 6d 61 6e 29                                   uman)
type.*func(*"".Human) t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 89 10 3d 38 00 04 04 36  ..........=8...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*func(*main.Human).+0
	rel 32+4 t=1 type.func(*"".Human)+0
type.func(*"".Human) t=8 dupok size=40
	0x0000 04 00 00 00 04 00 00 00 c5 ad 13 fa 02 04 04 33  ...............3
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 01 00 00 00 00 00 00 00                          ........
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*func(*main.Human).+0
	rel 28+4 t=6 type.*func(*"".Human)+0
	rel 36+4 t=1 type.*"".Human+0
type..namedata.Say. t=8 dupok size=6
	0x0000 01 00 03 53 61 79                                ...Say
type..namedata.*func(). t=8 dupok size=10
	0x0000 00 00 07 2a 66 75 6e 63 28 29                    ...*func()
type.*func() t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 9b 90 75 1b 00 04 04 36  ..........u....6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*func().+0
	rel 32+4 t=1 type.func()+0
type.func() t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 f6 bc 82 f6 02 04 04 33  ...............3
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*func().+0
	rel 28+4 t=6 type.*func()+0
type.*"".Human t=8 size=68
	0x0000 04 00 00 00 04 00 00 00 d3 86 b1 a0 01 04 04 36  ...............6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 01 00 00 00 10 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*main.Human.+0
	rel 28+4 t=6 type.**"".Human+0
	rel 32+4 t=1 type."".Human+0
	rel 36+4 t=5 type..importpath."".+0
	rel 52+4 t=5 type..namedata.Say.+0
	rel 56+4 t=24 type.func()+0
	rel 60+4 t=24 "".(*Human).Say+0
	rel 64+4 t=24 "".(*Human).Say+0
type..namedata.name. t=8 dupok size=7
	0x0000 00 00 04 6e 61 6d 65                             ...name
type."".Human t=8 size=76
	0x0000 08 00 00 00 04 00 00 00 2a 87 1d c5 07 04 04 19  ........*.......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 01 00 00 00 01 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 1c 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00              ............
	rel 16+4 t=1 runtime.algarray+56
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*main.Human.+0
	rel 28+4 t=5 type.*"".Human+0
	rel 32+4 t=1 type..importpath."".+0
	rel 36+4 t=1 type."".Human+64
	rel 48+4 t=5 type..importpath."".+0
	rel 64+4 t=1 type..namedata.name.+0
	rel 68+4 t=1 type.string+0
type..namedata.**main.Dog. t=8 dupok size=13
	0x0000 00 00 0a 2a 2a 6d 61 69 6e 2e 44 6f 67           ...**main.Dog
type.**"".Dog t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 61 82 50 4b 00 04 04 36  ........a.PK...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.**main.Dog.+0
	rel 32+4 t=1 type.*"".Dog+0
type..namedata.*main.Dog. t=8 dupok size=12
	0x0000 01 00 09 2a 6d 61 69 6e 2e 44 6f 67              ...*main.Dog
type..namedata.*func(*main.Dog). t=8 dupok size=19
	0x0000 00 00 10 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 44  ...*func(*main.D
	0x0010 6f 67 29                                         og)
type.*func(*"".Dog) t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 32 d1 84 3b 00 04 04 36  ........2..;...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*func(*main.Dog).+0
	rel 32+4 t=1 type.func(*"".Dog)+0
type.func(*"".Dog) t=8 dupok size=40
	0x0000 04 00 00 00 04 00 00 00 2e de 7a fb 02 04 04 33  ..........z....3
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 01 00 00 00 00 00 00 00                          ........
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*func(*main.Dog).+0
	rel 28+4 t=6 type.*func(*"".Dog)+0
	rel 36+4 t=1 type.*"".Dog+0
type.*"".Dog t=8 size=68
	0x0000 04 00 00 00 04 00 00 00 62 c0 18 19 01 04 04 36  ........b......6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 01 00 00 00 10 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*main.Dog.+0
	rel 28+4 t=6 type.**"".Dog+0
	rel 32+4 t=1 type."".Dog+0
	rel 36+4 t=5 type..importpath."".+0
	rel 52+4 t=5 type..namedata.Say.+0
	rel 56+4 t=24 type.func()+0
	rel 60+4 t=24 "".(*Dog).Say+0
	rel 64+4 t=24 "".(*Dog).Say+0
type."".Dog t=8 size=76
	0x0000 08 00 00 00 04 00 00 00 bf 36 61 92 07 04 04 19  .........6a.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 01 00 00 00 01 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 1c 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00              ............
	rel 16+4 t=1 runtime.algarray+56
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*main.Dog.+0
	rel 28+4 t=5 type.*"".Dog+0
	rel 32+4 t=1 type..importpath."".+0
	rel 36+4 t=1 type."".Dog+64
	rel 48+4 t=5 type..importpath."".+0
	rel 64+4 t=1 type..namedata.name.+0
	rel 68+4 t=1 type.string+0
type..namedata.*chan int. t=8 dupok size=12
	0x0000 00 00 09 2a 63 68 61 6e 20 69 6e 74              ...*chan int
type.*chan int t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 ed 7b ed 3b 00 04 04 36  .........{.;...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*chan int.+0
	rel 32+4 t=1 type.chan int+0
type.chan int t=8 dupok size=40
	0x0000 04 00 00 00 04 00 00 00 91 55 cb 71 02 04 04 32  .........U.q...2
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 03 00 00 00                          ........
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*chan int.+0
	rel 28+4 t=6 type.*chan int+0
	rel 32+4 t=1 type.int+0
type..namedata.*[]uint8. t=8 dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 a5 8e d0 69 00 04 04 36  ...........i...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]uint8.+0
	rel 32+4 t=1 type.[]uint8+0
type.[]uint8 t=8 dupok size=36
	0x0000 0c 00 00 00 04 00 00 00 df 7e 2e 38 02 04 04 17  .........~.8....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]uint8.+0
	rel 28+4 t=6 type.*[]uint8+0
	rel 32+4 t=1 type.uint8+0
runtime.gcbits. t=8 dupok size=0
type..namedata.*[4]uint8. t=8 dupok size=12
	0x0000 00 00 09 2a 5b 34 5d 75 69 6e 74 38              ...*[4]uint8
type.[4]uint8 t=8 dupok size=44
	0x0000 04 00 00 00 00 00 00 00 84 42 18 1c 02 01 01 91  .........B......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 04 00 00 00              ............
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.+0
	rel 24+4 t=5 type..namedata.*[4]uint8.+0
	rel 28+4 t=6 type.*[4]uint8+0
	rel 32+4 t=1 type.uint8+0
	rel 36+4 t=1 type.[]uint8+0
type.*[4]uint8 t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 47 b3 e3 a1 00 04 04 36  ........G......6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[4]uint8.+0
	rel 32+4 t=1 type.[4]uint8+0
type..namedata.*interface {}. t=8 dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 4f 0f 96 9d 00 04 04 36  ........O......6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*interface {}.+0
	rel 32+4 t=1 type.interface {}+0
runtime.gcbits.03 t=8 dupok size=1
	0x0000 03                                               .
type.interface {} t=8 dupok size=48
	0x0000 08 00 00 00 08 00 00 00 e7 57 a0 18 02 04 04 14  .........W......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 16+4 t=1 runtime.algarray+72
	rel 20+4 t=1 runtime.gcbits.03+0
	rel 24+4 t=5 type..namedata.*interface {}.+0
	rel 28+4 t=6 type.*interface {}+0
	rel 36+4 t=1 type.interface {}+48
type..namedata.*[]interface {}. t=8 dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 f3 04 9a e7 00 04 04 36  ...............6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]interface {}.+0
	rel 32+4 t=1 type.[]interface {}+0
type.[]interface {} t=8 dupok size=36
	0x0000 0c 00 00 00 04 00 00 00 70 93 ea 2f 02 04 04 17  ........p../....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]interface {}.+0
	rel 28+4 t=6 type.*[]interface {}+0
	rel 32+4 t=1 type.interface {}+0
type..namedata.*[1]interface {}. t=8 dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} t=8 dupok size=36
	0x0000 04 00 00 00 04 00 00 00 bf 03 a8 35 00 04 04 36  ...........5...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[1]interface {}.+0
	rel 32+4 t=1 type.[1]interface {}+0
type.[1]interface {} t=8 dupok size=44
	0x0000 08 00 00 00 08 00 00 00 50 91 5b fa 02 04 04 11  ........P.[.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 01 00 00 00              ............
	rel 16+4 t=1 runtime.algarray+72
	rel 20+4 t=1 runtime.gcbits.03+0
	rel 24+4 t=5 type..namedata.*[1]interface {}.+0
	rel 28+4 t=6 type.*[1]interface {}+0
	rel 32+4 t=1 type.interface {}+0
	rel 36+4 t=1 type.[]interface {}+0
go.info."".Speaker.Say t=45 dupok size=39
	0x0000 02 22 22 2e 53 70 65 61 6b 65 72 2e 53 61 79 00  ."".Speaker.Say.
	0x0010 00 00 00 00 00 00 00 00 01 05 2e 74 68 69 73 00  ...........this.
	0x0020 01 9c 00 00 00 00 00                             .......
	rel 16+4 t=1 "".Speaker.Say+0
	rel 20+4 t=1 "".Speaker.Say+65
	rel 34+4 t=28 go.info."".Speaker+0
type..namedata.*main.Speaker. t=8 dupok size=16
	0x0000 01 00 0d 2a 6d 61 69 6e 2e 53 70 65 61 6b 65 72  ...*main.Speaker
type.*"".Speaker t=8 size=36
	0x0000 04 00 00 00 04 00 00 00 fc 35 05 c8 00 04 04 36  .........5.....6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*main.Speaker.+0
	rel 32+4 t=1 type."".Speaker+0
type."".Speaker t=8 size=72
	0x0000 08 00 00 00 08 00 00 00 85 fb 88 82 07 04 04 14  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 01 00 00 00 01 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 16+4 t=1 runtime.algarray+64
	rel 20+4 t=1 runtime.gcbits.03+0
	rel 24+4 t=5 type..namedata.*main.Speaker.+0
	rel 28+4 t=5 type.*"".Speaker+0
	rel 32+4 t=1 type..importpath."".+0
	rel 36+4 t=1 type."".Speaker+64
	rel 48+4 t=5 type..importpath."".+0
	rel 64+4 t=5 type..namedata.Say.+0
	rel 68+4 t=5 type.func()+0
go.itab.*"".Human,"".Speaker t=32 dupok size=24
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+4 t=1 type."".Speaker+0
	rel 4+4 t=1 type.*"".Human+0
go.itablink.*"".Human,"".Speaker t=8 dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 go.itab.*"".Human,"".Speaker+0
go.itab.*"".Dog,"".Speaker t=32 dupok size=24
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+4 t=1 type."".Speaker+0
	rel 4+4 t=1 type.*"".Dog+0
go.itablink.*"".Dog,"".Speaker t=8 dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 go.itab.*"".Dog,"".Speaker+0
type..importpath.fmt. t=8 dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.unsafe. t=8 dupok size=9
	0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
