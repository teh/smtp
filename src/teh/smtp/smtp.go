// line 1 "smtp.ragel"
package smtp

import (
	"bytes"
	"errors"
	"fmt"
)

// Good introduction on smtp: http://cr.yp.to/smtp/request.html The
// email address parser is actually wrong, see
// http://en.wikipedia.org/wiki/Email_address#Local_part
// for a correct parser. It's an OK approximation for now though.

// line 98 "smtp.ragel"

// line 22 "smtp.go"
const smtp_start = 1
const smtp_first_final = 455
const smtp_error = 0

const smtp_en_main = 1

// line 101 "smtp.ragel"

var Dangling = errors.New("DANGLING")

func NewParser() *Parser {
	return &Parser{cs: smtp_en_main}
}

// Feed new bytes while keepting track of state in parser.current. If
// error is `Dangling` then more data is needed for a successfull
// parse. If error is nil then parser.current contains the new action
// item. Bytes from `data` are copied, enabling the called to recycle
// data slice.
//
// Gotchas: * current is not transactional and can contain
// intermediate values. * Parsing is not bounded: the internal
// parser.buffer grows as much as the input, restrictions need to be
// imposed from the outside (using e.g. a LimitReader).
func (parser *Parser) Feed(data []byte) (remaining []byte, err error) {
	var p, pb int
	cs := parser.cs
	pe := len(data)

	// line 54 "smtp.go"
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 35:
			goto st_case_35
		case 36:
			goto st_case_36
		case 37:
			goto st_case_37
		case 38:
			goto st_case_38
		case 39:
			goto st_case_39
		case 40:
			goto st_case_40
		case 41:
			goto st_case_41
		case 42:
			goto st_case_42
		case 43:
			goto st_case_43
		case 44:
			goto st_case_44
		case 45:
			goto st_case_45
		case 46:
			goto st_case_46
		case 455:
			goto st_case_455
		case 47:
			goto st_case_47
		case 48:
			goto st_case_48
		case 49:
			goto st_case_49
		case 50:
			goto st_case_50
		case 51:
			goto st_case_51
		case 52:
			goto st_case_52
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
		case 75:
			goto st_case_75
		case 76:
			goto st_case_76
		case 77:
			goto st_case_77
		case 78:
			goto st_case_78
		case 79:
			goto st_case_79
		case 80:
			goto st_case_80
		case 81:
			goto st_case_81
		case 82:
			goto st_case_82
		case 83:
			goto st_case_83
		case 84:
			goto st_case_84
		case 85:
			goto st_case_85
		case 86:
			goto st_case_86
		case 87:
			goto st_case_87
		case 88:
			goto st_case_88
		case 89:
			goto st_case_89
		case 90:
			goto st_case_90
		case 91:
			goto st_case_91
		case 92:
			goto st_case_92
		case 93:
			goto st_case_93
		case 94:
			goto st_case_94
		case 95:
			goto st_case_95
		case 96:
			goto st_case_96
		case 97:
			goto st_case_97
		case 98:
			goto st_case_98
		case 99:
			goto st_case_99
		case 100:
			goto st_case_100
		case 101:
			goto st_case_101
		case 102:
			goto st_case_102
		case 103:
			goto st_case_103
		case 104:
			goto st_case_104
		case 105:
			goto st_case_105
		case 106:
			goto st_case_106
		case 107:
			goto st_case_107
		case 108:
			goto st_case_108
		case 109:
			goto st_case_109
		case 110:
			goto st_case_110
		case 111:
			goto st_case_111
		case 112:
			goto st_case_112
		case 113:
			goto st_case_113
		case 114:
			goto st_case_114
		case 115:
			goto st_case_115
		case 116:
			goto st_case_116
		case 117:
			goto st_case_117
		case 118:
			goto st_case_118
		case 119:
			goto st_case_119
		case 120:
			goto st_case_120
		case 121:
			goto st_case_121
		case 122:
			goto st_case_122
		case 123:
			goto st_case_123
		case 124:
			goto st_case_124
		case 125:
			goto st_case_125
		case 126:
			goto st_case_126
		case 127:
			goto st_case_127
		case 128:
			goto st_case_128
		case 129:
			goto st_case_129
		case 130:
			goto st_case_130
		case 131:
			goto st_case_131
		case 132:
			goto st_case_132
		case 133:
			goto st_case_133
		case 134:
			goto st_case_134
		case 135:
			goto st_case_135
		case 136:
			goto st_case_136
		case 137:
			goto st_case_137
		case 138:
			goto st_case_138
		case 139:
			goto st_case_139
		case 140:
			goto st_case_140
		case 141:
			goto st_case_141
		case 142:
			goto st_case_142
		case 143:
			goto st_case_143
		case 144:
			goto st_case_144
		case 145:
			goto st_case_145
		case 146:
			goto st_case_146
		case 147:
			goto st_case_147
		case 148:
			goto st_case_148
		case 149:
			goto st_case_149
		case 150:
			goto st_case_150
		case 151:
			goto st_case_151
		case 152:
			goto st_case_152
		case 153:
			goto st_case_153
		case 154:
			goto st_case_154
		case 155:
			goto st_case_155
		case 156:
			goto st_case_156
		case 157:
			goto st_case_157
		case 158:
			goto st_case_158
		case 159:
			goto st_case_159
		case 160:
			goto st_case_160
		case 161:
			goto st_case_161
		case 162:
			goto st_case_162
		case 163:
			goto st_case_163
		case 164:
			goto st_case_164
		case 165:
			goto st_case_165
		case 166:
			goto st_case_166
		case 167:
			goto st_case_167
		case 168:
			goto st_case_168
		case 169:
			goto st_case_169
		case 170:
			goto st_case_170
		case 171:
			goto st_case_171
		case 172:
			goto st_case_172
		case 173:
			goto st_case_173
		case 174:
			goto st_case_174
		case 175:
			goto st_case_175
		case 176:
			goto st_case_176
		case 177:
			goto st_case_177
		case 178:
			goto st_case_178
		case 179:
			goto st_case_179
		case 180:
			goto st_case_180
		case 181:
			goto st_case_181
		case 182:
			goto st_case_182
		case 183:
			goto st_case_183
		case 184:
			goto st_case_184
		case 185:
			goto st_case_185
		case 186:
			goto st_case_186
		case 187:
			goto st_case_187
		case 188:
			goto st_case_188
		case 189:
			goto st_case_189
		case 190:
			goto st_case_190
		case 191:
			goto st_case_191
		case 192:
			goto st_case_192
		case 193:
			goto st_case_193
		case 194:
			goto st_case_194
		case 195:
			goto st_case_195
		case 196:
			goto st_case_196
		case 197:
			goto st_case_197
		case 198:
			goto st_case_198
		case 199:
			goto st_case_199
		case 200:
			goto st_case_200
		case 201:
			goto st_case_201
		case 202:
			goto st_case_202
		case 203:
			goto st_case_203
		case 204:
			goto st_case_204
		case 205:
			goto st_case_205
		case 206:
			goto st_case_206
		case 207:
			goto st_case_207
		case 208:
			goto st_case_208
		case 209:
			goto st_case_209
		case 210:
			goto st_case_210
		case 211:
			goto st_case_211
		case 212:
			goto st_case_212
		case 213:
			goto st_case_213
		case 214:
			goto st_case_214
		case 215:
			goto st_case_215
		case 216:
			goto st_case_216
		case 217:
			goto st_case_217
		case 218:
			goto st_case_218
		case 219:
			goto st_case_219
		case 220:
			goto st_case_220
		case 221:
			goto st_case_221
		case 222:
			goto st_case_222
		case 223:
			goto st_case_223
		case 224:
			goto st_case_224
		case 225:
			goto st_case_225
		case 226:
			goto st_case_226
		case 227:
			goto st_case_227
		case 228:
			goto st_case_228
		case 229:
			goto st_case_229
		case 230:
			goto st_case_230
		case 231:
			goto st_case_231
		case 232:
			goto st_case_232
		case 233:
			goto st_case_233
		case 234:
			goto st_case_234
		case 235:
			goto st_case_235
		case 236:
			goto st_case_236
		case 237:
			goto st_case_237
		case 238:
			goto st_case_238
		case 239:
			goto st_case_239
		case 240:
			goto st_case_240
		case 241:
			goto st_case_241
		case 242:
			goto st_case_242
		case 243:
			goto st_case_243
		case 244:
			goto st_case_244
		case 245:
			goto st_case_245
		case 246:
			goto st_case_246
		case 247:
			goto st_case_247
		case 248:
			goto st_case_248
		case 249:
			goto st_case_249
		case 250:
			goto st_case_250
		case 251:
			goto st_case_251
		case 252:
			goto st_case_252
		case 253:
			goto st_case_253
		case 254:
			goto st_case_254
		case 255:
			goto st_case_255
		case 256:
			goto st_case_256
		case 257:
			goto st_case_257
		case 258:
			goto st_case_258
		case 259:
			goto st_case_259
		case 260:
			goto st_case_260
		case 261:
			goto st_case_261
		case 262:
			goto st_case_262
		case 263:
			goto st_case_263
		case 264:
			goto st_case_264
		case 265:
			goto st_case_265
		case 266:
			goto st_case_266
		case 267:
			goto st_case_267
		case 268:
			goto st_case_268
		case 269:
			goto st_case_269
		case 270:
			goto st_case_270
		case 271:
			goto st_case_271
		case 272:
			goto st_case_272
		case 273:
			goto st_case_273
		case 274:
			goto st_case_274
		case 275:
			goto st_case_275
		case 276:
			goto st_case_276
		case 277:
			goto st_case_277
		case 278:
			goto st_case_278
		case 279:
			goto st_case_279
		case 280:
			goto st_case_280
		case 281:
			goto st_case_281
		case 282:
			goto st_case_282
		case 283:
			goto st_case_283
		case 284:
			goto st_case_284
		case 285:
			goto st_case_285
		case 286:
			goto st_case_286
		case 287:
			goto st_case_287
		case 288:
			goto st_case_288
		case 289:
			goto st_case_289
		case 290:
			goto st_case_290
		case 291:
			goto st_case_291
		case 292:
			goto st_case_292
		case 293:
			goto st_case_293
		case 294:
			goto st_case_294
		case 295:
			goto st_case_295
		case 296:
			goto st_case_296
		case 297:
			goto st_case_297
		case 298:
			goto st_case_298
		case 299:
			goto st_case_299
		case 300:
			goto st_case_300
		case 301:
			goto st_case_301
		case 302:
			goto st_case_302
		case 303:
			goto st_case_303
		case 304:
			goto st_case_304
		case 305:
			goto st_case_305
		case 306:
			goto st_case_306
		case 307:
			goto st_case_307
		case 308:
			goto st_case_308
		case 309:
			goto st_case_309
		case 310:
			goto st_case_310
		case 311:
			goto st_case_311
		case 312:
			goto st_case_312
		case 313:
			goto st_case_313
		case 314:
			goto st_case_314
		case 315:
			goto st_case_315
		case 316:
			goto st_case_316
		case 317:
			goto st_case_317
		case 318:
			goto st_case_318
		case 319:
			goto st_case_319
		case 320:
			goto st_case_320
		case 321:
			goto st_case_321
		case 322:
			goto st_case_322
		case 323:
			goto st_case_323
		case 324:
			goto st_case_324
		case 325:
			goto st_case_325
		case 326:
			goto st_case_326
		case 327:
			goto st_case_327
		case 328:
			goto st_case_328
		case 329:
			goto st_case_329
		case 330:
			goto st_case_330
		case 331:
			goto st_case_331
		case 332:
			goto st_case_332
		case 333:
			goto st_case_333
		case 334:
			goto st_case_334
		case 335:
			goto st_case_335
		case 336:
			goto st_case_336
		case 337:
			goto st_case_337
		case 338:
			goto st_case_338
		case 339:
			goto st_case_339
		case 340:
			goto st_case_340
		case 341:
			goto st_case_341
		case 342:
			goto st_case_342
		case 343:
			goto st_case_343
		case 344:
			goto st_case_344
		case 345:
			goto st_case_345
		case 346:
			goto st_case_346
		case 347:
			goto st_case_347
		case 348:
			goto st_case_348
		case 349:
			goto st_case_349
		case 350:
			goto st_case_350
		case 351:
			goto st_case_351
		case 352:
			goto st_case_352
		case 353:
			goto st_case_353
		case 354:
			goto st_case_354
		case 355:
			goto st_case_355
		case 356:
			goto st_case_356
		case 357:
			goto st_case_357
		case 358:
			goto st_case_358
		case 359:
			goto st_case_359
		case 360:
			goto st_case_360
		case 361:
			goto st_case_361
		case 362:
			goto st_case_362
		case 363:
			goto st_case_363
		case 364:
			goto st_case_364
		case 365:
			goto st_case_365
		case 366:
			goto st_case_366
		case 367:
			goto st_case_367
		case 368:
			goto st_case_368
		case 369:
			goto st_case_369
		case 370:
			goto st_case_370
		case 371:
			goto st_case_371
		case 372:
			goto st_case_372
		case 373:
			goto st_case_373
		case 456:
			goto st_case_456
		case 374:
			goto st_case_374
		case 375:
			goto st_case_375
		case 376:
			goto st_case_376
		case 377:
			goto st_case_377
		case 378:
			goto st_case_378
		case 379:
			goto st_case_379
		case 380:
			goto st_case_380
		case 381:
			goto st_case_381
		case 382:
			goto st_case_382
		case 383:
			goto st_case_383
		case 384:
			goto st_case_384
		case 385:
			goto st_case_385
		case 386:
			goto st_case_386
		case 387:
			goto st_case_387
		case 388:
			goto st_case_388
		case 389:
			goto st_case_389
		case 390:
			goto st_case_390
		case 391:
			goto st_case_391
		case 392:
			goto st_case_392
		case 393:
			goto st_case_393
		case 394:
			goto st_case_394
		case 395:
			goto st_case_395
		case 396:
			goto st_case_396
		case 397:
			goto st_case_397
		case 398:
			goto st_case_398
		case 399:
			goto st_case_399
		case 400:
			goto st_case_400
		case 401:
			goto st_case_401
		case 402:
			goto st_case_402
		case 403:
			goto st_case_403
		case 404:
			goto st_case_404
		case 405:
			goto st_case_405
		case 406:
			goto st_case_406
		case 407:
			goto st_case_407
		case 408:
			goto st_case_408
		case 409:
			goto st_case_409
		case 410:
			goto st_case_410
		case 411:
			goto st_case_411
		case 412:
			goto st_case_412
		case 413:
			goto st_case_413
		case 414:
			goto st_case_414
		case 415:
			goto st_case_415
		case 416:
			goto st_case_416
		case 417:
			goto st_case_417
		case 418:
			goto st_case_418
		case 419:
			goto st_case_419
		case 420:
			goto st_case_420
		case 421:
			goto st_case_421
		case 422:
			goto st_case_422
		case 423:
			goto st_case_423
		case 424:
			goto st_case_424
		case 425:
			goto st_case_425
		case 426:
			goto st_case_426
		case 427:
			goto st_case_427
		case 428:
			goto st_case_428
		case 429:
			goto st_case_429
		case 430:
			goto st_case_430
		case 431:
			goto st_case_431
		case 432:
			goto st_case_432
		case 433:
			goto st_case_433
		case 434:
			goto st_case_434
		case 457:
			goto st_case_457
		case 435:
			goto st_case_435
		case 436:
			goto st_case_436
		case 437:
			goto st_case_437
		case 438:
			goto st_case_438
		case 439:
			goto st_case_439
		case 440:
			goto st_case_440
		case 441:
			goto st_case_441
		case 442:
			goto st_case_442
		case 443:
			goto st_case_443
		case 444:
			goto st_case_444
		case 445:
			goto st_case_445
		case 446:
			goto st_case_446
		case 447:
			goto st_case_447
		case 448:
			goto st_case_448
		case 449:
			goto st_case_449
		case 450:
			goto st_case_450
		case 451:
			goto st_case_451
		case 452:
			goto st_case_452
		case 453:
			goto st_case_453
		case 454:
			goto st_case_454
		}
		goto st_end
	st_case_1:
		switch data[p] {
		case 43:
			goto ctr0
		case 61:
			goto ctr0
		case 65:
			goto ctr2
		case 68:
			goto ctr3
		case 69:
			goto ctr4
		case 72:
			goto ctr5
		case 77:
			goto ctr6
		case 81:
			goto ctr7
		case 82:
			goto ctr8
		case 83:
			goto ctr9
		case 86:
			goto ctr10
		}
		switch {
		case data[p] < 73:
			switch {
			case data[p] < 66:
				if 47 <= data[p] && data[p] <= 57 {
					goto ctr0
				}
			case data[p] > 67:
				if 70 <= data[p] && data[p] <= 71 {
					goto ctr0
				}
			default:
				goto ctr0
			}
		case data[p] > 76:
			switch {
			case data[p] < 84:
				if 78 <= data[p] && data[p] <= 80 {
					goto ctr0
				}
			case data[p] > 85:
				switch {
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto ctr0
					}
				case data[p] >= 87:
					goto ctr0
				}
			default:
				goto ctr0
			}
		default:
			goto ctr0
		}
		{
			goto st0

		}
	st_case_0:
	st0:
		cs = 0
		goto _out
	ctr0:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		// line 1059 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 43:
			goto st7
		case 61:
			goto st7
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st7
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st7
			}
		default:
			goto st7
		}
		{
			goto st0

		}
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		switch data[p] {
		case 43:
			goto st8
		case 61:
			goto st8
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		{
			goto st0

		}
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 43:
			goto st9
		case 61:
			goto st9
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st9
			}
		default:
			goto st9
		}
		{
			goto st0

		}
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 43:
			goto st10
		case 61:
			goto st10
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st10
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st10
			}
		default:
			goto st10
		}
		{
			goto st0

		}
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 43:
			goto st11
		case 61:
			goto st11
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st11
			}
		default:
			goto st11
		}
		{
			goto st0

		}
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 43:
			goto st12
		case 61:
			goto st12
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st12
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st12
			}
		default:
			goto st12
		}
		{
			goto st0

		}
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 43:
			goto st13
		case 61:
			goto st13
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st13
			}
		default:
			goto st13
		}
		{
			goto st0

		}
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 43:
			goto st14
		case 61:
			goto st14
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st14
			}
		default:
			goto st14
		}
		{
			goto st0

		}
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 43:
			goto st15
		case 61:
			goto st15
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st15
			}
		default:
			goto st15
		}
		{
			goto st0

		}
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 43:
			goto st16
		case 61:
			goto st16
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		{
			goto st0

		}
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 43:
			goto st17
		case 61:
			goto st17
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st17
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st17
			}
		default:
			goto st17
		}
		{
			goto st0

		}
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 43:
			goto st18
		case 61:
			goto st18
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st18
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st18
			}
		default:
			goto st18
		}
		{
			goto st0

		}
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 43:
			goto st19
		case 61:
			goto st19
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st19
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st19
			}
		default:
			goto st19
		}
		{
			goto st0

		}
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 43:
			goto st20
		case 61:
			goto st20
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st20
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st20
			}
		default:
			goto st20
		}
		{
			goto st0

		}
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		switch data[p] {
		case 43:
			goto st21
		case 61:
			goto st21
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st21
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st21
			}
		default:
			goto st21
		}
		{
			goto st0

		}
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 43:
			goto st22
		case 61:
			goto st22
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		default:
			goto st22
		}
		{
			goto st0

		}
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 43:
			goto st23
		case 61:
			goto st23
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		{
			goto st0

		}
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch data[p] {
		case 43:
			goto st24
		case 61:
			goto st24
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st24
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st24
			}
		default:
			goto st24
		}
		{
			goto st0

		}
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 43:
			goto st25
		case 61:
			goto st25
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st25
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st25
			}
		default:
			goto st25
		}
		{
			goto st0

		}
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 43:
			goto st26
		case 61:
			goto st26
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		{
			goto st0

		}
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 43:
			goto st27
		case 61:
			goto st27
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st27
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st27
			}
		default:
			goto st27
		}
		{
			goto st0

		}
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 43:
			goto st28
		case 61:
			goto st28
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st28
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st28
			}
		default:
			goto st28
		}
		{
			goto st0

		}
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 43:
			goto st29
		case 61:
			goto st29
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st29
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st29
			}
		default:
			goto st29
		}
		{
			goto st0

		}
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 43:
			goto st30
		case 61:
			goto st30
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st30
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st30
			}
		default:
			goto st30
		}
		{
			goto st0

		}
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 43:
			goto st31
		case 61:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st31
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st31
			}
		default:
			goto st31
		}
		{
			goto st0

		}
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 43:
			goto st32
		case 61:
			goto st32
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st32
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st32
			}
		default:
			goto st32
		}
		{
			goto st0

		}
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 43:
			goto st33
		case 61:
			goto st33
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st33
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st33
			}
		default:
			goto st33
		}
		{
			goto st0

		}
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 43:
			goto st34
		case 61:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st34
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st34
			}
		default:
			goto st34
		}
		{
			goto st0

		}
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 43:
			goto st35
		case 61:
			goto st35
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st35
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st35
			}
		default:
			goto st35
		}
		{
			goto st0

		}
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 43:
			goto st36
		case 61:
			goto st36
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st36
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st36
			}
		default:
			goto st36
		}
		{
			goto st0

		}
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 43:
			goto st37
		case 61:
			goto st37
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st37
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st37
			}
		default:
			goto st37
		}
		{
			goto st0

		}
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 43:
			goto st38
		case 61:
			goto st38
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st38
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st38
			}
		default:
			goto st38
		}
		{
			goto st0

		}
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 43:
			goto st39
		case 61:
			goto st39
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st39
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st39
			}
		default:
			goto st39
		}
		{
			goto st0

		}
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 43:
			goto st40
		case 61:
			goto st40
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st40
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st40
			}
		default:
			goto st40
		}
		{
			goto st0

		}
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 43:
			goto st41
		case 61:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st41
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		{
			goto st0

		}
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 43:
			goto st42
		case 61:
			goto st42
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		{
			goto st0

		}
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 43:
			goto st43
		case 61:
			goto st43
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st43
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st43
			}
		default:
			goto st43
		}
		{
			goto st0

		}
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 43:
			goto st44
		case 61:
			goto st44
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st44
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st44
			}
		default:
			goto st44
		}
		{
			goto st0

		}
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 43:
			goto st45
		case 61:
			goto st45
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st45
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st45
			}
		default:
			goto st45
		}
		{
			goto st0

		}
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 43:
			goto st46
		case 61:
			goto st46
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st46
			}
		default:
			goto st46
		}
		{
			goto st0

		}
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 10:
			goto ctr55
		case 13:
			goto ctr56
		case 43:
			goto st46
		case 61:
			goto st46
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st46
			}
		default:
			goto st46
		}
		{
			goto st0

		}
	ctr55:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 92 "smtp.ragel"
		{
			parser.current.Verb = VerbBASE64
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr57:
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr72:
		// line 89 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_CRAM_MD5
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr83:
		// line 88 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_DIGEST_MD5
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr89:
		// line 87 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_PLAIN
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr93:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 86 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_PLAIN
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr99:
		// line 84 "smtp.ragel"
		{
			parser.current.Verb = VerbDATA
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr107:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 80 "smtp.ragel"
		{
			parser.current.Verb = VerbEHLO
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr215:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 79 "smtp.ragel"
		{
			parser.current.Verb = VerbHELO
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr337:
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr351:
		// line 71 "smtp.ragel"
		{
			parser.current.BodyType = BodyType7BIT
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr361:
		// line 69 "smtp.ragel"
		{
			parser.current.BodyType = BodyType8BITMIME
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr451:
		// line 85 "smtp.ragel"
		{
			parser.current.Verb = VerbQUIT
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr472:
		// line 83 "smtp.ragel"
		{
			parser.current.Verb = VerbRCPT
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr499:
		// line 81 "smtp.ragel"
		{
			parser.current.Verb = VerbRSET
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr508:
		// line 90 "smtp.ragel"
		{
			parser.current.Verb = VerbSTARTTLS
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	ctr515:
		// line 91 "smtp.ragel"
		{
			parser.current.Verb = VerbVRFY
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 455
				goto _out
			}
		}
		goto st455
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		// line 2463 "smtp.go"
		{
			goto st0

		}
	ctr56:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 92 "smtp.ragel"
		{
			parser.current.Verb = VerbBASE64
		}
		goto st47
	ctr73:
		// line 89 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_CRAM_MD5
		}
		goto st47
	ctr84:
		// line 88 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_DIGEST_MD5
		}
		goto st47
	ctr90:
		// line 87 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_PLAIN
		}
		goto st47
	ctr94:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 86 "smtp.ragel"
		{
			parser.current.Verb = VerbAUTH_PLAIN
		}
		goto st47
	ctr100:
		// line 84 "smtp.ragel"
		{
			parser.current.Verb = VerbDATA
		}
		goto st47
	ctr108:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 80 "smtp.ragel"
		{
			parser.current.Verb = VerbEHLO
		}
		goto st47
	ctr216:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		// line 79 "smtp.ragel"
		{
			parser.current.Verb = VerbHELO
		}
		goto st47
	ctr338:
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		goto st47
	ctr352:
		// line 71 "smtp.ragel"
		{
			parser.current.BodyType = BodyType7BIT
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		goto st47
	ctr362:
		// line 69 "smtp.ragel"
		{
			parser.current.BodyType = BodyType8BITMIME
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		goto st47
	ctr452:
		// line 85 "smtp.ragel"
		{
			parser.current.Verb = VerbQUIT
		}
		goto st47
	ctr473:
		// line 83 "smtp.ragel"
		{
			parser.current.Verb = VerbRCPT
		}
		goto st47
	ctr500:
		// line 81 "smtp.ragel"
		{
			parser.current.Verb = VerbRSET
		}
		goto st47
	ctr509:
		// line 90 "smtp.ragel"
		{
			parser.current.Verb = VerbSTARTTLS
		}
		goto st47
	ctr516:
		// line 91 "smtp.ragel"
		{
			parser.current.Verb = VerbVRFY
		}
		goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		// line 2587 "smtp.go"
		if data[p] == 10 {
			goto ctr57
		}
		{
			goto st0

		}
	ctr2:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		// line 2608 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 85:
			goto st49
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 84:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 86:
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 84:
			goto st50
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 85:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 72:
			goto st51
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 71:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 73:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 32:
			goto st52
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 67:
			goto st53
		case 68:
			goto st61
		case 80:
			goto st71
		}
		{
			goto st0

		}
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		if data[p] == 82 {
			goto st54
		}
		{
			goto st0

		}
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		if data[p] == 65 {
			goto st55
		}
		{
			goto st0

		}
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		if data[p] == 77 {
			goto st56
		}
		{
			goto st0

		}
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		if data[p] == 45 {
			goto st57
		}
		{
			goto st0

		}
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		if data[p] == 77 {
			goto st58
		}
		{
			goto st0

		}
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		if data[p] == 68 {
			goto st59
		}
		{
			goto st0

		}
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 53 {
			goto st60
		}
		{
			goto st0

		}
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 10:
			goto ctr72
		case 13:
			goto ctr73
		}
		{
			goto st0

		}
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		if data[p] == 73 {
			goto st62
		}
		{
			goto st0

		}
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 71 {
			goto st63
		}
		{
			goto st0

		}
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		if data[p] == 69 {
			goto st64
		}
		{
			goto st0

		}
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		if data[p] == 83 {
			goto st65
		}
		{
			goto st0

		}
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 84 {
			goto st66
		}
		{
			goto st0

		}
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		if data[p] == 45 {
			goto st67
		}
		{
			goto st0

		}
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		if data[p] == 77 {
			goto st68
		}
		{
			goto st0

		}
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 68 {
			goto st69
		}
		{
			goto st0

		}
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		if data[p] == 53 {
			goto st70
		}
		{
			goto st0

		}
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 10:
			goto ctr83
		case 13:
			goto ctr84
		}
		{
			goto st0

		}
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		if data[p] == 76 {
			goto st72
		}
		{
			goto st0

		}
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 65 {
			goto st73
		}
		{
			goto st0

		}
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 73 {
			goto st74
		}
		{
			goto st0

		}
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		if data[p] == 78 {
			goto st75
		}
		{
			goto st0

		}
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 10:
			goto ctr89
		case 13:
			goto ctr90
		case 32:
			goto st76
		}
		{
			goto st0

		}
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 45:
			goto ctr92
		case 61:
			goto ctr92
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto ctr92
			}
		case data[p] >= 65:
			goto ctr92
		}
		{
			goto st0

		}
	ctr92:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		// line 3075 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr93
		case 13:
			goto ctr94
		case 45:
			goto st77
		case 61:
			goto st77
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st77
			}
		case data[p] >= 65:
			goto st77
		}
		{
			goto st0

		}
	ctr3:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		// line 3111 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 65:
			goto st79
		}
		switch {
		case data[p] < 66:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 84:
			goto st80
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 85:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 65:
			goto st81
		}
		switch {
		case data[p] < 66:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 10:
			goto ctr99
		case 13:
			goto ctr100
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	ctr4:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		// line 3243 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 72:
			goto st83
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 71:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 73:
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 76:
			goto st84
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 75:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 77:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 79:
			goto st85
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 78:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 80:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 32:
			goto st86
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		if data[p] == 91 {
			goto ctr106
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto ctr105
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto ctr105
			}
		default:
			goto ctr105
		}
		{
			goto st0

		}
	ctr105:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		// line 3407 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr107
		case 13:
			goto ctr108
		case 46:
			goto st88
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st87
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st87
			}
		default:
			goto st87
		}
		{
			goto st0

		}
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st87
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st87
			}
		default:
			goto st87
		}
		{
			goto st0

		}
	ctr106:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		// line 3466 "smtp.go"
		switch data[p] {
		case 48:
			goto st90
		case 73:
			goto st180
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st176
		}
		{
			goto st0

		}
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 46:
			goto st91
		case 58:
			goto st116
		case 120:
			goto st173
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st113
		}
		{
			goto st0

		}
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		if data[p] == 48 {
			goto st92
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st108
		}
		{
			goto st0

		}
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 46:
			goto st93
		case 93:
			goto st100
		case 120:
			goto st111
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st108
		}
		{
			goto st0

		}
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		if data[p] == 48 {
			goto st94
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st103
		}
		{
			goto st0

		}
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 46:
			goto st95
		case 93:
			goto st100
		case 120:
			goto st106
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st103
		}
		{
			goto st0

		}
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		if data[p] == 48 {
			goto st96
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st97
		}
		{
			goto st0

		}
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 93:
			goto st100
		case 120:
			goto st101
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st97
		}
		{
			goto st0

		}
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		if data[p] == 93 {
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st98
		}
		{
			goto st0

		}
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		if data[p] == 93 {
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st99
		}
		{
			goto st0

		}
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		if data[p] == 93 {
			goto st100
		}
		{
			goto st0

		}
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 10:
			goto ctr107
		case 13:
			goto ctr108
		}
		{
			goto st0

		}
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st102
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st102
			}
		default:
			goto st102
		}
		{
			goto st0

		}
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		if data[p] == 93 {
			goto st100
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st99
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st99
			}
		default:
			goto st99
		}
		{
			goto st0

		}
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 46:
			goto st95
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st104
		}
		{
			goto st0

		}
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 46:
			goto st95
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st105
		}
		{
			goto st0

		}
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 46:
			goto st95
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st107
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st107
			}
		default:
			goto st107
		}
		{
			goto st0

		}
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 46:
			goto st95
		case 93:
			goto st100
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st105
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st105
			}
		default:
			goto st105
		}
		{
			goto st0

		}
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 46:
			goto st93
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st109
		}
		{
			goto st0

		}
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 46:
			goto st93
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st110
		}
		{
			goto st0

		}
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 46:
			goto st93
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st112
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st112
			}
		default:
			goto st112
		}
		{
			goto st0

		}
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 46:
			goto st93
		case 93:
			goto st100
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st110
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st110
			}
		default:
			goto st110
		}
		{
			goto st0

		}
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 46:
			goto st91
		case 58:
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st114
		}
		{
			goto st0

		}
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 46:
			goto st91
		case 58:
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st115
		}
		{
			goto st0

		}
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 46:
			goto st91
		case 58:
			goto st116
		}
		{
			goto st0

		}
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		if data[p] == 58 {
			goto st172
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st117
		}
		{
			goto st0

		}
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 58:
			goto st121
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st118
		}
		{
			goto st0

		}
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 58:
			goto st121
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st119
		}
		{
			goto st0

		}
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 58:
			goto st121
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st120
		}
		{
			goto st0

		}
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		switch data[p] {
		case 58:
			goto st121
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
		if data[p] == 58 {
			goto st171
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		{
			goto st0

		}
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 58:
			goto st126
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st123
		}
		{
			goto st0

		}
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 58:
			goto st126
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st124
		}
		{
			goto st0

		}
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 58:
			goto st126
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st125
		}
		{
			goto st0

		}
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 58:
			goto st126
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		if data[p] == 58 {
			goto st170
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st127
		}
		{
			goto st0

		}
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 58:
			goto st131
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st128
		}
		{
			goto st0

		}
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 58:
			goto st131
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st129
		}
		{
			goto st0

		}
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 58:
			goto st131
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st130
		}
		{
			goto st0

		}
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 58:
			goto st131
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		if data[p] == 58 {
			goto st169
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st132
		}
		{
			goto st0

		}
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		switch data[p] {
		case 58:
			goto st136
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st133
		}
		{
			goto st0

		}
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 58:
			goto st136
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st134
		}
		{
			goto st0

		}
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 58:
			goto st136
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st135
		}
		{
			goto st0

		}
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 58:
			goto st136
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		if data[p] == 58 {
			goto st168
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st137
		}
		{
			goto st0

		}
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st138
		}
		{
			goto st0

		}
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st139
		}
		{
			goto st0

		}
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st140
		}
		{
			goto st0

		}
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 58:
			goto st141
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		if data[p] == 58 {
			goto st167
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st142
		}
		{
			goto st0

		}
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 58:
			goto st146
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st143
		}
		{
			goto st0

		}
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 58:
			goto st146
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st144
		}
		{
			goto st0

		}
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 58:
			goto st146
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st145
		}
		{
			goto st0

		}
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 58:
			goto st146
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		if data[p] == 58 {
			goto st166
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st147
		}
		{
			goto st0

		}
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 58:
			goto st151
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st148
		}
		{
			goto st0

		}
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 58:
			goto st151
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st149
		}
		{
			goto st0

		}
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
		switch data[p] {
		case 58:
			goto st151
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st150
		}
		{
			goto st0

		}
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		switch data[p] {
		case 58:
			goto st151
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		if data[p] == 58 {
			goto st165
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st152
		}
		{
			goto st0

		}
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st153
		}
		{
			goto st0

		}
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st154
		}
		{
			goto st0

		}
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st155
		}
		{
			goto st0

		}
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 58:
			goto st156
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		if data[p] == 58 {
			goto st164
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st157
		}
		{
			goto st0

		}
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 58:
			goto st161
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st158
		}
		{
			goto st0

		}
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 58:
			goto st161
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st159
		}
		{
			goto st0

		}
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 58:
			goto st161
		case 93:
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st160
		}
		{
			goto st0

		}
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 58:
			goto st161
		case 93:
			goto st100
		}
		{
			goto st0

		}
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		if data[p] == 58 {
			goto st163
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st162
		}
		{
			goto st0

		}
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		if data[p] == 93 {
			goto st100
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st97
		}
		{
			goto st0

		}
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		if 48 <= data[p] && data[p] <= 57 {
			goto st162
		}
		{
			goto st0

		}
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		if 48 <= data[p] && data[p] <= 57 {
			goto st157
		}
		{
			goto st0

		}
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		if 48 <= data[p] && data[p] <= 57 {
			goto st152
		}
		{
			goto st0

		}
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		if 48 <= data[p] && data[p] <= 57 {
			goto st147
		}
		{
			goto st0

		}
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if 48 <= data[p] && data[p] <= 57 {
			goto st142
		}
		{
			goto st0

		}
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		if 48 <= data[p] && data[p] <= 57 {
			goto st137
		}
		{
			goto st0

		}
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		if 48 <= data[p] && data[p] <= 57 {
			goto st132
		}
		{
			goto st0

		}
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		if 48 <= data[p] && data[p] <= 57 {
			goto st127
		}
		{
			goto st0

		}
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		if 48 <= data[p] && data[p] <= 57 {
			goto st122
		}
		{
			goto st0

		}
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		if 48 <= data[p] && data[p] <= 57 {
			goto st117
		}
		{
			goto st0

		}
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st174
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st174
			}
		default:
			goto st174
		}
		{
			goto st0

		}
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		if data[p] == 46 {
			goto st91
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st175
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st175
			}
		default:
			goto st175
		}
		{
			goto st0

		}
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		if data[p] == 46 {
			goto st91
		}
		{
			goto st0

		}
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 46:
			goto st91
		case 58:
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st177
		}
		{
			goto st0

		}
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 46:
			goto st91
		case 58:
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st178
		}
		{
			goto st0

		}
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 46:
			goto st91
		case 58:
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st179
		}
		{
			goto st0

		}
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		if data[p] == 58 {
			goto st116
		}
		{
			goto st0

		}
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		if data[p] == 80 {
			goto st181
		}
		{
			goto st0

		}
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		if data[p] == 118 {
			goto st182
		}
		{
			goto st0

		}
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		if data[p] == 54 {
			goto st183
		}
		{
			goto st0

		}
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		if data[p] == 58 {
			goto st184
		}
		{
			goto st0

		}
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		if 48 <= data[p] && data[p] <= 57 {
			goto st185
		}
		{
			goto st0

		}
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		if data[p] == 58 {
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st186
		}
		{
			goto st0

		}
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		if data[p] == 58 {
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st187
		}
		{
			goto st0

		}
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		if data[p] == 58 {
			goto st116
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st179
		}
		{
			goto st0

		}
	ctr5:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		// line 5101 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 69:
			goto st189
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 68:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 70:
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 76:
			goto st190
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 75:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 77:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 79:
			goto st191
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 78:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 80:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 32:
			goto st192
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		if data[p] == 91 {
			goto ctr214
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto ctr213
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto ctr213
			}
		default:
			goto ctr213
		}
		{
			goto st0

		}
	ctr213:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		// line 5265 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr215
		case 13:
			goto ctr216
		case 46:
			goto st194
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st193
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st193
			}
		default:
			goto st193
		}
		{
			goto st0

		}
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st193
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st193
			}
		default:
			goto st193
		}
		{
			goto st0

		}
	ctr214:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st195
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		// line 5324 "smtp.go"
		switch data[p] {
		case 48:
			goto st196
		case 73:
			goto st286
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st282
		}
		{
			goto st0

		}
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 46:
			goto st197
		case 58:
			goto st222
		case 120:
			goto st279
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st219
		}
		{
			goto st0

		}
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		if data[p] == 48 {
			goto st198
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st214
		}
		{
			goto st0

		}
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 46:
			goto st199
		case 93:
			goto st206
		case 120:
			goto st217
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st214
		}
		{
			goto st0

		}
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		if data[p] == 48 {
			goto st200
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st209
		}
		{
			goto st0

		}
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 46:
			goto st201
		case 93:
			goto st206
		case 120:
			goto st212
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st209
		}
		{
			goto st0

		}
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		if data[p] == 48 {
			goto st202
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st203
		}
		{
			goto st0

		}
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 93:
			goto st206
		case 120:
			goto st207
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st203
		}
		{
			goto st0

		}
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		if data[p] == 93 {
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st204
		}
		{
			goto st0

		}
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		if data[p] == 93 {
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st205
		}
		{
			goto st0

		}
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		if data[p] == 93 {
			goto st206
		}
		{
			goto st0

		}
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		switch data[p] {
		case 10:
			goto ctr215
		case 13:
			goto ctr216
		}
		{
			goto st0

		}
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st208
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st208
			}
		default:
			goto st208
		}
		{
			goto st0

		}
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		if data[p] == 93 {
			goto st206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st205
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st205
			}
		default:
			goto st205
		}
		{
			goto st0

		}
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		switch data[p] {
		case 46:
			goto st201
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st210
		}
		{
			goto st0

		}
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		switch data[p] {
		case 46:
			goto st201
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st211
		}
		{
			goto st0

		}
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 46:
			goto st201
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st213
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st213
			}
		default:
			goto st213
		}
		{
			goto st0

		}
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		switch data[p] {
		case 46:
			goto st201
		case 93:
			goto st206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st211
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st211
			}
		default:
			goto st211
		}
		{
			goto st0

		}
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		switch data[p] {
		case 46:
			goto st199
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st215
		}
		{
			goto st0

		}
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 46:
			goto st199
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st216
		}
		{
			goto st0

		}
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 46:
			goto st199
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st218
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st218
			}
		default:
			goto st218
		}
		{
			goto st0

		}
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 46:
			goto st199
		case 93:
			goto st206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st216
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st216
			}
		default:
			goto st216
		}
		{
			goto st0

		}
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 46:
			goto st197
		case 58:
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st220
		}
		{
			goto st0

		}
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 46:
			goto st197
		case 58:
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st221
		}
		{
			goto st0

		}
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 46:
			goto st197
		case 58:
			goto st222
		}
		{
			goto st0

		}
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		if data[p] == 58 {
			goto st278
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st223
		}
		{
			goto st0

		}
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 58:
			goto st227
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st224
		}
		{
			goto st0

		}
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 58:
			goto st227
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st225
		}
		{
			goto st0

		}
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 58:
			goto st227
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st226
		}
		{
			goto st0

		}
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 58:
			goto st227
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		if data[p] == 58 {
			goto st277
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st228
		}
		{
			goto st0

		}
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 58:
			goto st232
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st229
		}
		{
			goto st0

		}
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 58:
			goto st232
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st230
		}
		{
			goto st0

		}
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 58:
			goto st232
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st231
		}
		{
			goto st0

		}
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 58:
			goto st232
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		if data[p] == 58 {
			goto st276
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st233
		}
		{
			goto st0

		}
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		switch data[p] {
		case 58:
			goto st237
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st234
		}
		{
			goto st0

		}
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		switch data[p] {
		case 58:
			goto st237
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st235
		}
		{
			goto st0

		}
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		switch data[p] {
		case 58:
			goto st237
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st236
		}
		{
			goto st0

		}
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		switch data[p] {
		case 58:
			goto st237
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		if data[p] == 58 {
			goto st275
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st238
		}
		{
			goto st0

		}
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		switch data[p] {
		case 58:
			goto st242
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st239
		}
		{
			goto st0

		}
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		switch data[p] {
		case 58:
			goto st242
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st240
		}
		{
			goto st0

		}
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 58:
			goto st242
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st241
		}
		{
			goto st0

		}
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		switch data[p] {
		case 58:
			goto st242
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		if data[p] == 58 {
			goto st274
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st243
		}
		{
			goto st0

		}
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 58:
			goto st247
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st244
		}
		{
			goto st0

		}
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 58:
			goto st247
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st245
		}
		{
			goto st0

		}
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		switch data[p] {
		case 58:
			goto st247
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st246
		}
		{
			goto st0

		}
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		switch data[p] {
		case 58:
			goto st247
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		if data[p] == 58 {
			goto st273
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st248
		}
		{
			goto st0

		}
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 58:
			goto st252
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st249
		}
		{
			goto st0

		}
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		switch data[p] {
		case 58:
			goto st252
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st250
		}
		{
			goto st0

		}
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 58:
			goto st252
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st251
		}
		{
			goto st0

		}
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 58:
			goto st252
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		if data[p] == 58 {
			goto st272
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st253
		}
		{
			goto st0

		}
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		switch data[p] {
		case 58:
			goto st257
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st254
		}
		{
			goto st0

		}
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		switch data[p] {
		case 58:
			goto st257
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st255
		}
		{
			goto st0

		}
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 58:
			goto st257
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st256
		}
		{
			goto st0

		}
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 58:
			goto st257
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		if data[p] == 58 {
			goto st271
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st258
		}
		{
			goto st0

		}
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		switch data[p] {
		case 58:
			goto st262
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st259
		}
		{
			goto st0

		}
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		switch data[p] {
		case 58:
			goto st262
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st260
		}
		{
			goto st0

		}
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 58:
			goto st262
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st261
		}
		{
			goto st0

		}
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 58:
			goto st262
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		if data[p] == 58 {
			goto st270
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st263
		}
		{
			goto st0

		}
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		switch data[p] {
		case 58:
			goto st267
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st264
		}
		{
			goto st0

		}
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 58:
			goto st267
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st265
		}
		{
			goto st0

		}
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		switch data[p] {
		case 58:
			goto st267
		case 93:
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st266
		}
		{
			goto st0

		}
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		switch data[p] {
		case 58:
			goto st267
		case 93:
			goto st206
		}
		{
			goto st0

		}
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		if data[p] == 58 {
			goto st269
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st268
		}
		{
			goto st0

		}
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		if data[p] == 93 {
			goto st206
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st203
		}
		{
			goto st0

		}
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		if 48 <= data[p] && data[p] <= 57 {
			goto st268
		}
		{
			goto st0

		}
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		if 48 <= data[p] && data[p] <= 57 {
			goto st263
		}
		{
			goto st0

		}
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		if 48 <= data[p] && data[p] <= 57 {
			goto st258
		}
		{
			goto st0

		}
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		if 48 <= data[p] && data[p] <= 57 {
			goto st253
		}
		{
			goto st0

		}
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		if 48 <= data[p] && data[p] <= 57 {
			goto st248
		}
		{
			goto st0

		}
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		if 48 <= data[p] && data[p] <= 57 {
			goto st243
		}
		{
			goto st0

		}
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		if 48 <= data[p] && data[p] <= 57 {
			goto st238
		}
		{
			goto st0

		}
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		if 48 <= data[p] && data[p] <= 57 {
			goto st233
		}
		{
			goto st0

		}
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		if 48 <= data[p] && data[p] <= 57 {
			goto st228
		}
		{
			goto st0

		}
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		if 48 <= data[p] && data[p] <= 57 {
			goto st223
		}
		{
			goto st0

		}
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st280
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st280
			}
		default:
			goto st280
		}
		{
			goto st0

		}
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		if data[p] == 46 {
			goto st197
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st281
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st281
			}
		default:
			goto st281
		}
		{
			goto st0

		}
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		if data[p] == 46 {
			goto st197
		}
		{
			goto st0

		}
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 46:
			goto st197
		case 58:
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st283
		}
		{
			goto st0

		}
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		switch data[p] {
		case 46:
			goto st197
		case 58:
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st284
		}
		{
			goto st0

		}
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 46:
			goto st197
		case 58:
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st285
		}
		{
			goto st0

		}
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		if data[p] == 58 {
			goto st222
		}
		{
			goto st0

		}
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		if data[p] == 80 {
			goto st287
		}
		{
			goto st0

		}
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		if data[p] == 118 {
			goto st288
		}
		{
			goto st0

		}
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		if data[p] == 54 {
			goto st289
		}
		{
			goto st0

		}
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		if data[p] == 58 {
			goto st290
		}
		{
			goto st0

		}
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		if 48 <= data[p] && data[p] <= 57 {
			goto st291
		}
		{
			goto st0

		}
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		if data[p] == 58 {
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st292
		}
		{
			goto st0

		}
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		if data[p] == 58 {
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st293
		}
		{
			goto st0

		}
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		if data[p] == 58 {
			goto st222
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st285
		}
		{
			goto st0

		}
	ctr6:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st294
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		// line 6959 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 65:
			goto st295
		}
		switch {
		case data[p] < 66:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 73:
			goto st296
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 72:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 74:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 76:
			goto st297
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 75:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 77:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		switch data[p] {
		case 32:
			goto st298
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 70:
			goto st299
		case 102:
			goto st299
		}
		{
			goto st0

		}
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 82:
			goto st300
		case 114:
			goto st300
		}
		{
			goto st0

		}
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 79:
			goto st301
		case 111:
			goto st301
		}
		{
			goto st0

		}
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		switch data[p] {
		case 77:
			goto st302
		case 109:
			goto st302
		}
		{
			goto st0

		}
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		if data[p] == 58 {
			goto st303
		}
		{
			goto st0

		}
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		if data[p] == 60 {
			goto st304
		}
		{
			goto st0

		}
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 33:
			goto ctr327
		case 34:
			goto ctr328
		case 64:
			goto ctr329
		case 92:
			goto ctr330
		}
		switch {
		case data[p] < 65:
			if 35 <= data[p] && data[p] <= 63 {
				goto ctr327
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr327
			}
		default:
			goto ctr327
		}
		{
			goto st0

		}
	ctr327:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st305
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		// line 7209 "smtp.go"
		switch data[p] {
		case 33:
			goto st305
		case 34:
			goto st306
		case 62:
			goto ctr333
		case 92:
			goto st334
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st305
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st305
			}
		default:
			goto st305
		}
		{
			goto st0

		}
	ctr328:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st306
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		// line 7249 "smtp.go"
		if data[p] == 92 {
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		switch data[p] {
		case 34:
			goto st305
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		if data[p] == 34 {
			goto st307
		}
		{
			goto st0

		}
	ctr333:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st309
	ctr413:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st309
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		// line 7337 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr337
		case 13:
			goto ctr338
		case 32:
			goto st310
		case 33:
			goto st305
		case 34:
			goto st306
		case 62:
			goto ctr333
		case 92:
			goto st334
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st305
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st305
			}
		default:
			goto st305
		}
		{
			goto st0

		}
	ctr353:
		// line 71 "smtp.ragel"
		{
			parser.current.BodyType = BodyType7BIT
		}
		goto st310
	ctr363:
		// line 69 "smtp.ragel"
		{
			parser.current.BodyType = BodyType8BITMIME
		}
		goto st310
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		// line 7385 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr337
		case 13:
			goto ctr338
		case 32:
			goto st310
		case 66:
			goto st311
		case 83:
			goto st328
		}
		{
			goto st0

		}
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		if data[p] == 79 {
			goto st312
		}
		{
			goto st0

		}
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		if data[p] == 68 {
			goto st313
		}
		{
			goto st0

		}
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		if data[p] == 89 {
			goto st314
		}
		{
			goto st0

		}
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		if data[p] == 61 {
			goto st315
		}
		{
			goto st0

		}
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 55:
			goto st316
		case 56:
			goto st320
		}
		{
			goto st0

		}
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		if data[p] == 66 {
			goto st317
		}
		{
			goto st0

		}
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		if data[p] == 73 {
			goto st318
		}
		{
			goto st0

		}
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		if data[p] == 84 {
			goto st319
		}
		{
			goto st0

		}
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 10:
			goto ctr351
		case 13:
			goto ctr352
		case 32:
			goto ctr353
		}
		{
			goto st0

		}
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		if data[p] == 66 {
			goto st321
		}
		{
			goto st0

		}
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		if data[p] == 73 {
			goto st322
		}
		{
			goto st0

		}
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		if data[p] == 84 {
			goto st323
		}
		{
			goto st0

		}
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		if data[p] == 77 {
			goto st324
		}
		{
			goto st0

		}
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		if data[p] == 73 {
			goto st325
		}
		{
			goto st0

		}
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		if data[p] == 77 {
			goto st326
		}
		{
			goto st0

		}
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		if data[p] == 69 {
			goto st327
		}
		{
			goto st0

		}
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		switch data[p] {
		case 10:
			goto ctr361
		case 13:
			goto ctr362
		case 32:
			goto ctr363
		}
		{
			goto st0

		}
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		if data[p] == 73 {
			goto st329
		}
		{
			goto st0

		}
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		if data[p] == 90 {
			goto st330
		}
		{
			goto st0

		}
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		if data[p] == 69 {
			goto st331
		}
		{
			goto st0

		}
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		if data[p] == 61 {
			goto st332
		}
		{
			goto st0

		}
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		if 48 <= data[p] && data[p] <= 57 {
			goto st333
		}
		{
			goto st0

		}
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 10:
			goto ctr337
		case 13:
			goto ctr338
		case 32:
			goto st310
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st333
		}
		{
			goto st0

		}
	ctr330:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st334
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		// line 7712 "smtp.go"
		if data[p] == 92 {
			goto st335
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st305
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st305
			}
		default:
			goto st305
		}
		{
			goto st0

		}
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		if data[p] == 34 {
			goto st305
		}
		{
			goto st0

		}
	ctr329:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st336
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		// line 7757 "smtp.go"
		switch data[p] {
		case 33:
			goto st339
		case 34:
			goto st340
		case 58:
			goto st305
		case 62:
			goto ctr373
		case 92:
			goto st399
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 32:
				if 35 <= data[p] && data[p] <= 57 {
					goto st339
				}
			default:
				goto st337
			}
		case data[p] > 61:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st339
				}
			case data[p] >= 63:
				goto st339
			}
		default:
			goto st339
		}
		{
			goto st337

		}
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		if data[p] == 58 {
			goto st338
		}
		if data[p] <= 57 {
			goto st337
		}
		{
			goto st337

		}
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		switch data[p] {
		case 33:
			goto ctr327
		case 34:
			goto ctr328
		case 92:
			goto ctr330
		}
		switch {
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr327
			}
		case data[p] >= 35:
			goto ctr327
		}
		{
			goto st0

		}
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		switch data[p] {
		case 33:
			goto st339
		case 34:
			goto st340
		case 58:
			goto st372
		case 62:
			goto ctr373
		case 92:
			goto st399
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 32:
				if 35 <= data[p] && data[p] <= 57 {
					goto st339
				}
			default:
				goto st337
			}
		case data[p] > 61:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st339
				}
			case data[p] >= 63:
				goto st339
			}
		default:
			goto st339
		}
		{
			goto st337

		}
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		switch data[p] {
		case 34:
			goto st337
		case 58:
			goto st342
		case 92:
			goto st371
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] > 31:
				if data[p] <= 33 {
					goto st341
				}
			default:
				goto st337
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st341
				}
			case data[p] >= 59:
				goto st341
			}
		default:
			goto st341
		}
		{
			goto st337

		}
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		switch data[p] {
		case 34:
			goto st339
		case 58:
			goto st342
		case 92:
			goto st371
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] > 31:
				if data[p] <= 33 {
					goto st341
				}
			default:
				goto st337
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st341
				}
			case data[p] >= 59:
				goto st341
			}
		default:
			goto st341
		}
		{
			goto st337

		}
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		switch data[p] {
		case 32:
			goto st307
		case 33:
			goto ctr380
		case 34:
			goto ctr381
		case 92:
			goto ctr382
		}
		switch {
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr380
			}
		case data[p] >= 35:
			goto ctr380
		}
		{
			goto st0

		}
	ctr380:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		// line 7997 "smtp.go"
		switch data[p] {
		case 32:
			goto st307
		case 33:
			goto st343
		case 34:
			goto st344
		case 62:
			goto ctr385
		case 92:
			goto st370
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st343
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st343
			}
		default:
			goto st343
		}
		{
			goto st0

		}
	ctr381:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		// line 8039 "smtp.go"
		switch data[p] {
		case 32:
			goto st307
		case 33:
			goto st343
		case 34:
			goto st306
		case 62:
			goto ctr385
		case 92:
			goto st370
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st343
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st343
			}
		default:
			goto st343
		}
		{
			goto st0

		}
	ctr385:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st345
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		// line 8082 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr337
		case 13:
			goto ctr338
		case 32:
			goto st346
		case 33:
			goto st343
		case 34:
			goto st344
		case 62:
			goto ctr385
		case 92:
			goto st370
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st343
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st343
			}
		default:
			goto st343
		}
		{
			goto st0

		}
	ctr399:
		// line 71 "smtp.ragel"
		{
			parser.current.BodyType = BodyType7BIT
		}
		goto st346
	ctr407:
		// line 69 "smtp.ragel"
		{
			parser.current.BodyType = BodyType8BITMIME
		}
		goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		// line 8130 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr337
		case 13:
			goto ctr338
		case 32:
			goto st346
		case 33:
			goto st307
		case 34:
			goto st305
		case 66:
			goto st347
		case 83:
			goto st364
		case 92:
			goto st308
		}
		switch {
		case data[p] < 67:
			if 35 <= data[p] && data[p] <= 65 {
				goto st307
			}
		case data[p] > 82:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 84:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 34:
			goto st305
		case 79:
			goto st348
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 78:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 80:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 34:
			goto st305
		case 68:
			goto st349
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 67:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 69:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 34:
			goto st305
		case 89:
			goto st350
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 88:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 90:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 34:
			goto st305
		case 61:
			goto st351
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 60:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 62:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		switch data[p] {
		case 34:
			goto st305
		case 55:
			goto st352
		case 56:
			goto st356
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 54:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 57:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		switch data[p] {
		case 34:
			goto st305
		case 66:
			goto st353
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 65:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 67:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		switch data[p] {
		case 34:
			goto st305
		case 73:
			goto st354
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 72:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 74:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		switch data[p] {
		case 34:
			goto st305
		case 84:
			goto st355
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 83:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 85:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 10:
			goto ctr351
		case 13:
			goto ctr352
		case 32:
			goto ctr399
		case 33:
			goto st307
		case 34:
			goto st305
		case 92:
			goto st308
		}
		switch {
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st307
			}
		case data[p] >= 35:
			goto st307
		}
		{
			goto st0

		}
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		switch data[p] {
		case 34:
			goto st305
		case 66:
			goto st357
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 65:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 67:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		switch data[p] {
		case 34:
			goto st305
		case 73:
			goto st358
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 72:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 74:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		switch data[p] {
		case 34:
			goto st305
		case 84:
			goto st359
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 83:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 85:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		switch data[p] {
		case 34:
			goto st305
		case 77:
			goto st360
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 76:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 78:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		switch data[p] {
		case 34:
			goto st305
		case 73:
			goto st361
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 72:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 74:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		switch data[p] {
		case 34:
			goto st305
		case 77:
			goto st362
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 76:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 78:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 34:
			goto st305
		case 69:
			goto st363
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 68:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 70:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 10:
			goto ctr361
		case 13:
			goto ctr362
		case 32:
			goto ctr407
		case 33:
			goto st307
		case 34:
			goto st305
		case 92:
			goto st308
		}
		switch {
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st307
			}
		case data[p] >= 35:
			goto st307
		}
		{
			goto st0

		}
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 34:
			goto st305
		case 73:
			goto st365
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 72:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 74:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 34:
			goto st305
		case 90:
			goto st366
		case 91:
			goto st307
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 89:
			if 93 <= data[p] && data[p] <= 126 {
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 34:
			goto st305
		case 69:
			goto st367
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 68:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 70:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		switch data[p] {
		case 34:
			goto st305
		case 61:
			goto st368
		case 92:
			goto st308
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st307
			}
		case data[p] > 60:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			case data[p] >= 62:
				goto st307
			}
		default:
			goto st307
		}
		{
			goto st0

		}
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 34:
			goto st305
		case 92:
			goto st308
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 33:
				if 35 <= data[p] {
					goto st307
				}
			case data[p] >= 32:
				goto st307
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			default:
				goto st307
			}
		default:
			goto st369
		}
		{
			goto st0

		}
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		switch data[p] {
		case 10:
			goto ctr337
		case 13:
			goto ctr338
		case 32:
			goto st346
		case 33:
			goto st307
		case 34:
			goto st305
		case 92:
			goto st308
		}
		switch {
		case data[p] < 48:
			if 35 <= data[p] {
				goto st307
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st307
				}
			default:
				goto st307
			}
		default:
			goto st369
		}
		{
			goto st0

		}
	ctr382:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st370
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		// line 8967 "smtp.go"
		switch data[p] {
		case 34:
			goto st307
		case 92:
			goto st335
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st305
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st305
			}
		default:
			goto st305
		}
		{
			goto st0

		}
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		switch data[p] {
		case 34:
			goto st341
		case 58:
			goto st338
		}
		switch {
		case data[p] > 33:
			if 35 <= data[p] && data[p] <= 57 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		switch data[p] {
		case 33:
			goto ctr327
		case 34:
			goto ctr328
		case 62:
			goto ctr413
		case 92:
			goto ctr330
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto ctr327
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr327
			}
		default:
			goto ctr327
		}
		{
			goto st0

		}
	ctr373:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st373
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		// line 9058 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr414
		case 13:
			goto ctr415
		case 32:
			goto st375
		case 33:
			goto st339
		case 34:
			goto st340
		case 58:
			goto st372
		case 62:
			goto ctr373
		case 92:
			goto st399
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] < 11:
				if data[p] <= 9 {
					goto st337
				}
			case data[p] > 12:
				if 14 <= data[p] && data[p] <= 31 {
					goto st337
				}
			default:
				goto st337
			}
		case data[p] > 57:
			switch {
			case data[p] < 63:
				if 59 <= data[p] && data[p] <= 61 {
					goto st339
				}
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st339
				}
			default:
				goto st339
			}
		default:
			goto st339
		}
		{
			goto st337

		}
	ctr417:
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 456
				goto _out
			}
		}
		goto st456
	ctr414:
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 456
				goto _out
			}
		}
		goto st456
	ctr429:
		// line 71 "smtp.ragel"
		{
			parser.current.BodyType = BodyType7BIT
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 456
				goto _out
			}
		}
		goto st456
	ctr439:
		// line 69 "smtp.ragel"
		{
			parser.current.BodyType = BodyType8BITMIME
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 456
				goto _out
			}
		}
		goto st456
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		// line 9155 "smtp.go"
		if data[p] == 58 {
			goto st338
		}
		if data[p] <= 57 {
			goto st337
		}
		{
			goto st337

		}
	ctr415:
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		goto st374
	ctr430:
		// line 71 "smtp.ragel"
		{
			parser.current.BodyType = BodyType7BIT
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		goto st374
	ctr440:
		// line 69 "smtp.ragel"
		{
			parser.current.BodyType = BodyType8BITMIME
		}
		// line 82 "smtp.ragel"
		{
			parser.current.Verb = VerbMAIL
		}
		goto st374
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		// line 9192 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr417
		case 58:
			goto st338
		}
		switch {
		case data[p] > 9:
			if 11 <= data[p] && data[p] <= 57 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	ctr431:
		// line 71 "smtp.ragel"
		{
			parser.current.BodyType = BodyType7BIT
		}
		goto st375
	ctr441:
		// line 69 "smtp.ragel"
		{
			parser.current.BodyType = BodyType8BITMIME
		}
		goto st375
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		// line 9226 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr414
		case 13:
			goto ctr415
		case 32:
			goto st375
		case 58:
			goto st338
		case 66:
			goto st376
		case 83:
			goto st393
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 9:
				if 11 <= data[p] && data[p] <= 12 {
					goto st337
				}
			default:
				goto st337
			}
		case data[p] > 31:
			switch {
			case data[p] < 59:
				if 33 <= data[p] && data[p] <= 57 {
					goto st337
				}
			case data[p] > 65:
				if 67 <= data[p] && data[p] <= 82 {
					goto st337
				}
			default:
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		switch data[p] {
		case 58:
			goto st338
		case 79:
			goto st377
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 78 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 58:
			goto st338
		case 68:
			goto st378
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 67 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		switch data[p] {
		case 58:
			goto st338
		case 89:
			goto st379
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 88 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		switch data[p] {
		case 58:
			goto st338
		case 61:
			goto st380
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 60 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		switch data[p] {
		case 55:
			goto st381
		case 56:
			goto st385
		case 57:
			goto st337
		case 58:
			goto st338
		}
		if data[p] <= 54 {
			goto st337
		}
		{
			goto st337

		}
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 58:
			goto st338
		case 66:
			goto st382
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 65 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 58:
			goto st338
		case 73:
			goto st383
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 72 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 58:
			goto st338
		case 84:
			goto st384
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 83 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		switch data[p] {
		case 10:
			goto ctr429
		case 13:
			goto ctr430
		case 32:
			goto ctr431
		case 58:
			goto st338
		}
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto st337
			}
		case data[p] > 12:
			switch {
			case data[p] > 31:
				if 33 <= data[p] && data[p] <= 57 {
					goto st337
				}
			case data[p] >= 14:
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 58:
			goto st338
		case 66:
			goto st386
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 65 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		switch data[p] {
		case 58:
			goto st338
		case 73:
			goto st387
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 72 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		switch data[p] {
		case 58:
			goto st338
		case 84:
			goto st388
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 83 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		switch data[p] {
		case 58:
			goto st338
		case 77:
			goto st389
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 76 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 58:
			goto st338
		case 73:
			goto st390
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 72 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		switch data[p] {
		case 58:
			goto st338
		case 77:
			goto st391
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 76 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		switch data[p] {
		case 58:
			goto st338
		case 69:
			goto st392
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 68 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 10:
			goto ctr439
		case 13:
			goto ctr440
		case 32:
			goto ctr441
		case 58:
			goto st338
		}
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto st337
			}
		case data[p] > 12:
			switch {
			case data[p] > 31:
				if 33 <= data[p] && data[p] <= 57 {
					goto st337
				}
			case data[p] >= 14:
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		switch data[p] {
		case 58:
			goto st338
		case 73:
			goto st394
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 72 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		switch data[p] {
		case 58:
			goto st338
		case 90:
			goto st395
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 89 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		switch data[p] {
		case 58:
			goto st338
		case 69:
			goto st396
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 68 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		switch data[p] {
		case 58:
			goto st338
		case 61:
			goto st397
		}
		switch {
		case data[p] > 57:
			if 59 <= data[p] && data[p] <= 60 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		if data[p] == 58 {
			goto st338
		}
		switch {
		case data[p] > 47:
			if data[p] <= 57 {
				goto st398
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		switch data[p] {
		case 10:
			goto ctr414
		case 13:
			goto ctr415
		case 32:
			goto st375
		case 58:
			goto st338
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 9:
				if 11 <= data[p] && data[p] <= 12 {
					goto st337
				}
			default:
				goto st337
			}
		case data[p] > 31:
			switch {
			case data[p] > 47:
				if data[p] <= 57 {
					goto st398
				}
			case data[p] >= 33:
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		switch data[p] {
		case 34:
			goto st337
		case 58:
			goto st372
		case 92:
			goto st400
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] > 31:
				if data[p] <= 33 {
					goto st339
				}
			default:
				goto st337
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st339
				}
			case data[p] >= 59:
				goto st339
			}
		default:
			goto st339
		}
		{
			goto st337

		}
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 34:
			goto st339
		case 58:
			goto st338
		}
		switch {
		case data[p] > 33:
			if 35 <= data[p] && data[p] <= 57 {
				goto st337
			}
		default:
			goto st337
		}
		{
			goto st337

		}
	ctr7:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st401
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		// line 9915 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 85:
			goto st402
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 84:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 86:
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 73:
			goto st403
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 72:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 74:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 84:
			goto st404
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 85:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 10:
			goto ctr451
		case 13:
			goto ctr452
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	ctr8:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st405
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		// line 10057 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 67:
			goto st406
		case 83:
			goto st438
		}
		switch {
		case data[p] < 68:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 66 {
					goto st3
				}
			case data[p] >= 47:
				goto st3
			}
		case data[p] > 82:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 84:
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 80:
			goto st407
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 79:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 81:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 84:
			goto st408
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 85:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 32:
			goto st409
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		switch data[p] {
		case 84:
			goto st410
		case 116:
			goto st410
		}
		{
			goto st0

		}
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		switch data[p] {
		case 79:
			goto st411
		case 111:
			goto st411
		}
		{
			goto st0

		}
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		if data[p] == 58 {
			goto st412
		}
		{
			goto st0

		}
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		if data[p] == 60 {
			goto st413
		}
		{
			goto st0

		}
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		switch data[p] {
		case 33:
			goto ctr462
		case 34:
			goto ctr463
		case 64:
			goto ctr464
		case 92:
			goto ctr465
		}
		switch {
		case data[p] < 65:
			if 35 <= data[p] && data[p] <= 63 {
				goto ctr462
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr462
			}
		default:
			goto ctr462
		}
		{
			goto st0

		}
	ctr462:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		// line 10289 "smtp.go"
		switch data[p] {
		case 33:
			goto st414
		case 34:
			goto st415
		case 62:
			goto ctr468
		case 92:
			goto st419
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st414
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st414
			}
		default:
			goto st414
		}
		{
			goto st0

		}
	ctr463:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st415
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		// line 10329 "smtp.go"
		if data[p] == 92 {
			goto st417
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st416
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st416
			}
		default:
			goto st416
		}
		{
			goto st0

		}
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 34:
			goto st414
		case 92:
			goto st417
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st416
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st416
			}
		default:
			goto st416
		}
		{
			goto st0

		}
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		if data[p] == 34 {
			goto st416
		}
		{
			goto st0

		}
	ctr468:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st418
	ctr492:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		// line 10417 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr472
		case 13:
			goto ctr473
		case 33:
			goto st414
		case 34:
			goto st415
		case 62:
			goto ctr468
		case 92:
			goto st419
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st414
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st414
			}
		default:
			goto st414
		}
		{
			goto st0

		}
	ctr465:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st419
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		// line 10461 "smtp.go"
		if data[p] == 92 {
			goto st420
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st414
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st414
			}
		default:
			goto st414
		}
		{
			goto st0

		}
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		if data[p] == 34 {
			goto st414
		}
		{
			goto st0

		}
	ctr464:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st421
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		// line 10506 "smtp.go"
		switch data[p] {
		case 33:
			goto st424
		case 34:
			goto st425
		case 58:
			goto st414
		case 62:
			goto ctr478
		case 92:
			goto st436
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 32:
				if 35 <= data[p] && data[p] <= 57 {
					goto st424
				}
			default:
				goto st422
			}
		case data[p] > 61:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st424
				}
			case data[p] >= 63:
				goto st424
			}
		default:
			goto st424
		}
		{
			goto st422

		}
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		if data[p] == 58 {
			goto st423
		}
		if data[p] <= 57 {
			goto st422
		}
		{
			goto st422

		}
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		switch data[p] {
		case 33:
			goto ctr462
		case 34:
			goto ctr463
		case 92:
			goto ctr465
		}
		switch {
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr462
			}
		case data[p] >= 35:
			goto ctr462
		}
		{
			goto st0

		}
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 33:
			goto st424
		case 34:
			goto st425
		case 58:
			goto st433
		case 62:
			goto ctr478
		case 92:
			goto st436
		}
		switch {
		case data[p] < 59:
			switch {
			case data[p] > 32:
				if 35 <= data[p] && data[p] <= 57 {
					goto st424
				}
			default:
				goto st422
			}
		case data[p] > 61:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st424
				}
			case data[p] >= 63:
				goto st424
			}
		default:
			goto st424
		}
		{
			goto st422

		}
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		switch data[p] {
		case 34:
			goto st422
		case 58:
			goto st427
		case 92:
			goto st432
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] > 31:
				if data[p] <= 33 {
					goto st426
				}
			default:
				goto st422
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st426
				}
			case data[p] >= 59:
				goto st426
			}
		default:
			goto st426
		}
		{
			goto st422

		}
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
		switch data[p] {
		case 34:
			goto st424
		case 58:
			goto st427
		case 92:
			goto st432
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] > 31:
				if data[p] <= 33 {
					goto st426
				}
			default:
				goto st422
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st426
				}
			case data[p] >= 59:
				goto st426
			}
		default:
			goto st426
		}
		{
			goto st422

		}
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 32:
			goto st416
		case 33:
			goto ctr485
		case 34:
			goto ctr486
		case 92:
			goto ctr487
		}
		switch {
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr485
			}
		case data[p] >= 35:
			goto ctr485
		}
		{
			goto st0

		}
	ctr485:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		// line 10746 "smtp.go"
		switch data[p] {
		case 32:
			goto st416
		case 33:
			goto st428
		case 34:
			goto st429
		case 62:
			goto ctr490
		case 92:
			goto st431
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st428
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st428
			}
		default:
			goto st428
		}
		{
			goto st0

		}
	ctr486:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		// line 10788 "smtp.go"
		switch data[p] {
		case 32:
			goto st416
		case 33:
			goto st428
		case 34:
			goto st415
		case 62:
			goto ctr490
		case 92:
			goto st431
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st428
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st428
			}
		default:
			goto st428
		}
		{
			goto st0

		}
	ctr490:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		// line 10831 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr472
		case 13:
			goto ctr473
		case 32:
			goto st416
		case 33:
			goto st428
		case 34:
			goto st429
		case 62:
			goto ctr490
		case 92:
			goto st431
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto st428
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st428
			}
		default:
			goto st428
		}
		{
			goto st0

		}
	ctr487:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st431
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		// line 10877 "smtp.go"
		switch data[p] {
		case 34:
			goto st416
		case 92:
			goto st420
		}
		switch {
		case data[p] < 35:
			if 32 <= data[p] && data[p] <= 33 {
				goto st414
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto st414
			}
		default:
			goto st414
		}
		{
			goto st0

		}
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		switch data[p] {
		case 34:
			goto st426
		case 58:
			goto st423
		}
		switch {
		case data[p] > 33:
			if 35 <= data[p] && data[p] <= 57 {
				goto st422
			}
		default:
			goto st422
		}
		{
			goto st422

		}
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 33:
			goto ctr462
		case 34:
			goto ctr463
		case 62:
			goto ctr492
		case 92:
			goto ctr465
		}
		switch {
		case data[p] < 63:
			if 35 <= data[p] && data[p] <= 61 {
				goto ctr462
			}
		case data[p] > 91:
			if 93 <= data[p] && data[p] <= 126 {
				goto ctr462
			}
		default:
			goto ctr462
		}
		{
			goto st0

		}
	ctr478:
		// line 26 "smtp.ragel"
		{

			parser.buffer.Write(data[pb:p])
			parser.current.Data = parser.buffer.Bytes()
			parser.buffer = nil
		}
		goto st434
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		// line 10968 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr493
		case 13:
			goto ctr494
		case 33:
			goto st424
		case 34:
			goto st425
		case 58:
			goto st433
		case 62:
			goto ctr478
		case 92:
			goto st436
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] < 11:
				if data[p] <= 9 {
					goto st422
				}
			case data[p] > 12:
				if 14 <= data[p] && data[p] <= 32 {
					goto st422
				}
			default:
				goto st422
			}
		case data[p] > 57:
			switch {
			case data[p] < 63:
				if 59 <= data[p] && data[p] <= 61 {
					goto st424
				}
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st424
				}
			default:
				goto st424
			}
		default:
			goto st424
		}
		{
			goto st422

		}
	ctr495:
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 457
				goto _out
			}
		}
		goto st457
	ctr493:
		// line 83 "smtp.ragel"
		{
			parser.current.Verb = VerbRCPT
		}
		// line 97 "smtp.ragel"
		{
			{
				p++
				cs = 457
				goto _out
			}
		}
		goto st457
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		// line 11039 "smtp.go"
		if data[p] == 58 {
			goto st423
		}
		if data[p] <= 57 {
			goto st422
		}
		{
			goto st422

		}
	ctr494:
		// line 83 "smtp.ragel"
		{
			parser.current.Verb = VerbRCPT
		}
		goto st435
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		// line 11060 "smtp.go"
		switch data[p] {
		case 10:
			goto ctr495
		case 58:
			goto st423
		}
		switch {
		case data[p] > 9:
			if 11 <= data[p] && data[p] <= 57 {
				goto st422
			}
		default:
			goto st422
		}
		{
			goto st422

		}
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		switch data[p] {
		case 34:
			goto st422
		case 58:
			goto st433
		case 92:
			goto st437
		}
		switch {
		case data[p] < 35:
			switch {
			case data[p] > 31:
				if data[p] <= 33 {
					goto st424
				}
			default:
				goto st422
			}
		case data[p] > 57:
			switch {
			case data[p] > 91:
				if 93 <= data[p] && data[p] <= 126 {
					goto st424
				}
			case data[p] >= 59:
				goto st424
			}
		default:
			goto st424
		}
		{
			goto st422

		}
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 34:
			goto st424
		case 58:
			goto st423
		}
		switch {
		case data[p] > 33:
			if 35 <= data[p] && data[p] <= 57 {
				goto st422
			}
		default:
			goto st422
		}
		{
			goto st422

		}
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 69:
			goto st439
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 68:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 70:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 84:
			goto st440
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 85:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		switch data[p] {
		case 10:
			goto ctr499
		case 13:
			goto ctr500
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	ctr9:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		// line 11253 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 84:
			goto st442
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 85:
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 65:
			goto st443
		}
		switch {
		case data[p] < 66:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 82:
			goto st444
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 81:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st5
				}
			case data[p] >= 83:
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		switch data[p] {
		case 43:
			goto st6
		case 61:
			goto st6
		case 84:
			goto st445
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st6
				}
			case data[p] >= 85:
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		switch data[p] {
		case 43:
			goto st7
		case 61:
			goto st7
		case 84:
			goto st446
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st7
			}
		case data[p] > 83:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st7
				}
			case data[p] >= 85:
				goto st7
			}
		default:
			goto st7
		}
		{
			goto st0

		}
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 43:
			goto st8
		case 61:
			goto st8
		case 76:
			goto st447
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] > 75:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st8
				}
			case data[p] >= 77:
				goto st8
			}
		default:
			goto st8
		}
		{
			goto st0

		}
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		switch data[p] {
		case 43:
			goto st9
		case 61:
			goto st9
		case 83:
			goto st448
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st9
			}
		case data[p] > 82:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st9
				}
			case data[p] >= 84:
				goto st9
			}
		default:
			goto st9
		}
		{
			goto st0

		}
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		switch data[p] {
		case 10:
			goto ctr508
		case 13:
			goto ctr509
		case 43:
			goto st10
		case 61:
			goto st10
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st10
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st10
			}
		default:
			goto st10
		}
		{
			goto st0

		}
	ctr10:
		// line 32 "smtp.ragel"
		{

			pb = p
			parser.buffer = bytes.NewBuffer(nil)
		}
		goto st449
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		// line 11526 "smtp.go"
		switch data[p] {
		case 43:
			goto st3
		case 61:
			goto st3
		case 82:
			goto st450
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st3
			}
		case data[p] > 81:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st3
				}
			case data[p] >= 83:
				goto st3
			}
		default:
			goto st3
		}
		{
			goto st0

		}
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		switch data[p] {
		case 43:
			goto st4
		case 61:
			goto st4
		case 70:
			goto st451
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 69:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st4
				}
			case data[p] >= 71:
				goto st4
			}
		default:
			goto st4
		}
		{
			goto st0

		}
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		switch data[p] {
		case 43:
			goto st5
		case 61:
			goto st5
		case 89:
			goto st452
		case 90:
			goto st5
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 88:
			if 97 <= data[p] && data[p] <= 122 {
				goto st5
			}
		default:
			goto st5
		}
		{
			goto st0

		}
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 32:
			goto st453
		case 43:
			goto st6
		case 61:
			goto st6
		}
		switch {
		case data[p] < 65:
			if 47 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st6
			}
		default:
			goto st6
		}
		{
			goto st0

		}
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		switch data[p] {
		case 10:
			goto st0
		case 13:
			goto st0
		}
		switch {
		case data[p] > 9:
			if 11 <= data[p] && data[p] <= 12 {
				goto st454
			}
		default:
			goto st454
		}
		{
			goto st454

		}
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 10:
			goto ctr515
		case 13:
			goto ctr516
		}
		switch {
		case data[p] > 9:
			if 11 <= data[p] && data[p] <= 12 {
				goto st454
			}
		default:
			goto st454
		}
		{
			goto st454

		}
	st_end:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof455:
		cs = 455
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
		goto _test_eof
	_test_eof79:
		cs = 79
		goto _test_eof
	_test_eof80:
		cs = 80
		goto _test_eof
	_test_eof81:
		cs = 81
		goto _test_eof
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof86:
		cs = 86
		goto _test_eof
	_test_eof87:
		cs = 87
		goto _test_eof
	_test_eof88:
		cs = 88
		goto _test_eof
	_test_eof89:
		cs = 89
		goto _test_eof
	_test_eof90:
		cs = 90
		goto _test_eof
	_test_eof91:
		cs = 91
		goto _test_eof
	_test_eof92:
		cs = 92
		goto _test_eof
	_test_eof93:
		cs = 93
		goto _test_eof
	_test_eof94:
		cs = 94
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof101:
		cs = 101
		goto _test_eof
	_test_eof102:
		cs = 102
		goto _test_eof
	_test_eof103:
		cs = 103
		goto _test_eof
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof106:
		cs = 106
		goto _test_eof
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof108:
		cs = 108
		goto _test_eof
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof114:
		cs = 114
		goto _test_eof
	_test_eof115:
		cs = 115
		goto _test_eof
	_test_eof116:
		cs = 116
		goto _test_eof
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof
	_test_eof130:
		cs = 130
		goto _test_eof
	_test_eof131:
		cs = 131
		goto _test_eof
	_test_eof132:
		cs = 132
		goto _test_eof
	_test_eof133:
		cs = 133
		goto _test_eof
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof135:
		cs = 135
		goto _test_eof
	_test_eof136:
		cs = 136
		goto _test_eof
	_test_eof137:
		cs = 137
		goto _test_eof
	_test_eof138:
		cs = 138
		goto _test_eof
	_test_eof139:
		cs = 139
		goto _test_eof
	_test_eof140:
		cs = 140
		goto _test_eof
	_test_eof141:
		cs = 141
		goto _test_eof
	_test_eof142:
		cs = 142
		goto _test_eof
	_test_eof143:
		cs = 143
		goto _test_eof
	_test_eof144:
		cs = 144
		goto _test_eof
	_test_eof145:
		cs = 145
		goto _test_eof
	_test_eof146:
		cs = 146
		goto _test_eof
	_test_eof147:
		cs = 147
		goto _test_eof
	_test_eof148:
		cs = 148
		goto _test_eof
	_test_eof149:
		cs = 149
		goto _test_eof
	_test_eof150:
		cs = 150
		goto _test_eof
	_test_eof151:
		cs = 151
		goto _test_eof
	_test_eof152:
		cs = 152
		goto _test_eof
	_test_eof153:
		cs = 153
		goto _test_eof
	_test_eof154:
		cs = 154
		goto _test_eof
	_test_eof155:
		cs = 155
		goto _test_eof
	_test_eof156:
		cs = 156
		goto _test_eof
	_test_eof157:
		cs = 157
		goto _test_eof
	_test_eof158:
		cs = 158
		goto _test_eof
	_test_eof159:
		cs = 159
		goto _test_eof
	_test_eof160:
		cs = 160
		goto _test_eof
	_test_eof161:
		cs = 161
		goto _test_eof
	_test_eof162:
		cs = 162
		goto _test_eof
	_test_eof163:
		cs = 163
		goto _test_eof
	_test_eof164:
		cs = 164
		goto _test_eof
	_test_eof165:
		cs = 165
		goto _test_eof
	_test_eof166:
		cs = 166
		goto _test_eof
	_test_eof167:
		cs = 167
		goto _test_eof
	_test_eof168:
		cs = 168
		goto _test_eof
	_test_eof169:
		cs = 169
		goto _test_eof
	_test_eof170:
		cs = 170
		goto _test_eof
	_test_eof171:
		cs = 171
		goto _test_eof
	_test_eof172:
		cs = 172
		goto _test_eof
	_test_eof173:
		cs = 173
		goto _test_eof
	_test_eof174:
		cs = 174
		goto _test_eof
	_test_eof175:
		cs = 175
		goto _test_eof
	_test_eof176:
		cs = 176
		goto _test_eof
	_test_eof177:
		cs = 177
		goto _test_eof
	_test_eof178:
		cs = 178
		goto _test_eof
	_test_eof179:
		cs = 179
		goto _test_eof
	_test_eof180:
		cs = 180
		goto _test_eof
	_test_eof181:
		cs = 181
		goto _test_eof
	_test_eof182:
		cs = 182
		goto _test_eof
	_test_eof183:
		cs = 183
		goto _test_eof
	_test_eof184:
		cs = 184
		goto _test_eof
	_test_eof185:
		cs = 185
		goto _test_eof
	_test_eof186:
		cs = 186
		goto _test_eof
	_test_eof187:
		cs = 187
		goto _test_eof
	_test_eof188:
		cs = 188
		goto _test_eof
	_test_eof189:
		cs = 189
		goto _test_eof
	_test_eof190:
		cs = 190
		goto _test_eof
	_test_eof191:
		cs = 191
		goto _test_eof
	_test_eof192:
		cs = 192
		goto _test_eof
	_test_eof193:
		cs = 193
		goto _test_eof
	_test_eof194:
		cs = 194
		goto _test_eof
	_test_eof195:
		cs = 195
		goto _test_eof
	_test_eof196:
		cs = 196
		goto _test_eof
	_test_eof197:
		cs = 197
		goto _test_eof
	_test_eof198:
		cs = 198
		goto _test_eof
	_test_eof199:
		cs = 199
		goto _test_eof
	_test_eof200:
		cs = 200
		goto _test_eof
	_test_eof201:
		cs = 201
		goto _test_eof
	_test_eof202:
		cs = 202
		goto _test_eof
	_test_eof203:
		cs = 203
		goto _test_eof
	_test_eof204:
		cs = 204
		goto _test_eof
	_test_eof205:
		cs = 205
		goto _test_eof
	_test_eof206:
		cs = 206
		goto _test_eof
	_test_eof207:
		cs = 207
		goto _test_eof
	_test_eof208:
		cs = 208
		goto _test_eof
	_test_eof209:
		cs = 209
		goto _test_eof
	_test_eof210:
		cs = 210
		goto _test_eof
	_test_eof211:
		cs = 211
		goto _test_eof
	_test_eof212:
		cs = 212
		goto _test_eof
	_test_eof213:
		cs = 213
		goto _test_eof
	_test_eof214:
		cs = 214
		goto _test_eof
	_test_eof215:
		cs = 215
		goto _test_eof
	_test_eof216:
		cs = 216
		goto _test_eof
	_test_eof217:
		cs = 217
		goto _test_eof
	_test_eof218:
		cs = 218
		goto _test_eof
	_test_eof219:
		cs = 219
		goto _test_eof
	_test_eof220:
		cs = 220
		goto _test_eof
	_test_eof221:
		cs = 221
		goto _test_eof
	_test_eof222:
		cs = 222
		goto _test_eof
	_test_eof223:
		cs = 223
		goto _test_eof
	_test_eof224:
		cs = 224
		goto _test_eof
	_test_eof225:
		cs = 225
		goto _test_eof
	_test_eof226:
		cs = 226
		goto _test_eof
	_test_eof227:
		cs = 227
		goto _test_eof
	_test_eof228:
		cs = 228
		goto _test_eof
	_test_eof229:
		cs = 229
		goto _test_eof
	_test_eof230:
		cs = 230
		goto _test_eof
	_test_eof231:
		cs = 231
		goto _test_eof
	_test_eof232:
		cs = 232
		goto _test_eof
	_test_eof233:
		cs = 233
		goto _test_eof
	_test_eof234:
		cs = 234
		goto _test_eof
	_test_eof235:
		cs = 235
		goto _test_eof
	_test_eof236:
		cs = 236
		goto _test_eof
	_test_eof237:
		cs = 237
		goto _test_eof
	_test_eof238:
		cs = 238
		goto _test_eof
	_test_eof239:
		cs = 239
		goto _test_eof
	_test_eof240:
		cs = 240
		goto _test_eof
	_test_eof241:
		cs = 241
		goto _test_eof
	_test_eof242:
		cs = 242
		goto _test_eof
	_test_eof243:
		cs = 243
		goto _test_eof
	_test_eof244:
		cs = 244
		goto _test_eof
	_test_eof245:
		cs = 245
		goto _test_eof
	_test_eof246:
		cs = 246
		goto _test_eof
	_test_eof247:
		cs = 247
		goto _test_eof
	_test_eof248:
		cs = 248
		goto _test_eof
	_test_eof249:
		cs = 249
		goto _test_eof
	_test_eof250:
		cs = 250
		goto _test_eof
	_test_eof251:
		cs = 251
		goto _test_eof
	_test_eof252:
		cs = 252
		goto _test_eof
	_test_eof253:
		cs = 253
		goto _test_eof
	_test_eof254:
		cs = 254
		goto _test_eof
	_test_eof255:
		cs = 255
		goto _test_eof
	_test_eof256:
		cs = 256
		goto _test_eof
	_test_eof257:
		cs = 257
		goto _test_eof
	_test_eof258:
		cs = 258
		goto _test_eof
	_test_eof259:
		cs = 259
		goto _test_eof
	_test_eof260:
		cs = 260
		goto _test_eof
	_test_eof261:
		cs = 261
		goto _test_eof
	_test_eof262:
		cs = 262
		goto _test_eof
	_test_eof263:
		cs = 263
		goto _test_eof
	_test_eof264:
		cs = 264
		goto _test_eof
	_test_eof265:
		cs = 265
		goto _test_eof
	_test_eof266:
		cs = 266
		goto _test_eof
	_test_eof267:
		cs = 267
		goto _test_eof
	_test_eof268:
		cs = 268
		goto _test_eof
	_test_eof269:
		cs = 269
		goto _test_eof
	_test_eof270:
		cs = 270
		goto _test_eof
	_test_eof271:
		cs = 271
		goto _test_eof
	_test_eof272:
		cs = 272
		goto _test_eof
	_test_eof273:
		cs = 273
		goto _test_eof
	_test_eof274:
		cs = 274
		goto _test_eof
	_test_eof275:
		cs = 275
		goto _test_eof
	_test_eof276:
		cs = 276
		goto _test_eof
	_test_eof277:
		cs = 277
		goto _test_eof
	_test_eof278:
		cs = 278
		goto _test_eof
	_test_eof279:
		cs = 279
		goto _test_eof
	_test_eof280:
		cs = 280
		goto _test_eof
	_test_eof281:
		cs = 281
		goto _test_eof
	_test_eof282:
		cs = 282
		goto _test_eof
	_test_eof283:
		cs = 283
		goto _test_eof
	_test_eof284:
		cs = 284
		goto _test_eof
	_test_eof285:
		cs = 285
		goto _test_eof
	_test_eof286:
		cs = 286
		goto _test_eof
	_test_eof287:
		cs = 287
		goto _test_eof
	_test_eof288:
		cs = 288
		goto _test_eof
	_test_eof289:
		cs = 289
		goto _test_eof
	_test_eof290:
		cs = 290
		goto _test_eof
	_test_eof291:
		cs = 291
		goto _test_eof
	_test_eof292:
		cs = 292
		goto _test_eof
	_test_eof293:
		cs = 293
		goto _test_eof
	_test_eof294:
		cs = 294
		goto _test_eof
	_test_eof295:
		cs = 295
		goto _test_eof
	_test_eof296:
		cs = 296
		goto _test_eof
	_test_eof297:
		cs = 297
		goto _test_eof
	_test_eof298:
		cs = 298
		goto _test_eof
	_test_eof299:
		cs = 299
		goto _test_eof
	_test_eof300:
		cs = 300
		goto _test_eof
	_test_eof301:
		cs = 301
		goto _test_eof
	_test_eof302:
		cs = 302
		goto _test_eof
	_test_eof303:
		cs = 303
		goto _test_eof
	_test_eof304:
		cs = 304
		goto _test_eof
	_test_eof305:
		cs = 305
		goto _test_eof
	_test_eof306:
		cs = 306
		goto _test_eof
	_test_eof307:
		cs = 307
		goto _test_eof
	_test_eof308:
		cs = 308
		goto _test_eof
	_test_eof309:
		cs = 309
		goto _test_eof
	_test_eof310:
		cs = 310
		goto _test_eof
	_test_eof311:
		cs = 311
		goto _test_eof
	_test_eof312:
		cs = 312
		goto _test_eof
	_test_eof313:
		cs = 313
		goto _test_eof
	_test_eof314:
		cs = 314
		goto _test_eof
	_test_eof315:
		cs = 315
		goto _test_eof
	_test_eof316:
		cs = 316
		goto _test_eof
	_test_eof317:
		cs = 317
		goto _test_eof
	_test_eof318:
		cs = 318
		goto _test_eof
	_test_eof319:
		cs = 319
		goto _test_eof
	_test_eof320:
		cs = 320
		goto _test_eof
	_test_eof321:
		cs = 321
		goto _test_eof
	_test_eof322:
		cs = 322
		goto _test_eof
	_test_eof323:
		cs = 323
		goto _test_eof
	_test_eof324:
		cs = 324
		goto _test_eof
	_test_eof325:
		cs = 325
		goto _test_eof
	_test_eof326:
		cs = 326
		goto _test_eof
	_test_eof327:
		cs = 327
		goto _test_eof
	_test_eof328:
		cs = 328
		goto _test_eof
	_test_eof329:
		cs = 329
		goto _test_eof
	_test_eof330:
		cs = 330
		goto _test_eof
	_test_eof331:
		cs = 331
		goto _test_eof
	_test_eof332:
		cs = 332
		goto _test_eof
	_test_eof333:
		cs = 333
		goto _test_eof
	_test_eof334:
		cs = 334
		goto _test_eof
	_test_eof335:
		cs = 335
		goto _test_eof
	_test_eof336:
		cs = 336
		goto _test_eof
	_test_eof337:
		cs = 337
		goto _test_eof
	_test_eof338:
		cs = 338
		goto _test_eof
	_test_eof339:
		cs = 339
		goto _test_eof
	_test_eof340:
		cs = 340
		goto _test_eof
	_test_eof341:
		cs = 341
		goto _test_eof
	_test_eof342:
		cs = 342
		goto _test_eof
	_test_eof343:
		cs = 343
		goto _test_eof
	_test_eof344:
		cs = 344
		goto _test_eof
	_test_eof345:
		cs = 345
		goto _test_eof
	_test_eof346:
		cs = 346
		goto _test_eof
	_test_eof347:
		cs = 347
		goto _test_eof
	_test_eof348:
		cs = 348
		goto _test_eof
	_test_eof349:
		cs = 349
		goto _test_eof
	_test_eof350:
		cs = 350
		goto _test_eof
	_test_eof351:
		cs = 351
		goto _test_eof
	_test_eof352:
		cs = 352
		goto _test_eof
	_test_eof353:
		cs = 353
		goto _test_eof
	_test_eof354:
		cs = 354
		goto _test_eof
	_test_eof355:
		cs = 355
		goto _test_eof
	_test_eof356:
		cs = 356
		goto _test_eof
	_test_eof357:
		cs = 357
		goto _test_eof
	_test_eof358:
		cs = 358
		goto _test_eof
	_test_eof359:
		cs = 359
		goto _test_eof
	_test_eof360:
		cs = 360
		goto _test_eof
	_test_eof361:
		cs = 361
		goto _test_eof
	_test_eof362:
		cs = 362
		goto _test_eof
	_test_eof363:
		cs = 363
		goto _test_eof
	_test_eof364:
		cs = 364
		goto _test_eof
	_test_eof365:
		cs = 365
		goto _test_eof
	_test_eof366:
		cs = 366
		goto _test_eof
	_test_eof367:
		cs = 367
		goto _test_eof
	_test_eof368:
		cs = 368
		goto _test_eof
	_test_eof369:
		cs = 369
		goto _test_eof
	_test_eof370:
		cs = 370
		goto _test_eof
	_test_eof371:
		cs = 371
		goto _test_eof
	_test_eof372:
		cs = 372
		goto _test_eof
	_test_eof373:
		cs = 373
		goto _test_eof
	_test_eof456:
		cs = 456
		goto _test_eof
	_test_eof374:
		cs = 374
		goto _test_eof
	_test_eof375:
		cs = 375
		goto _test_eof
	_test_eof376:
		cs = 376
		goto _test_eof
	_test_eof377:
		cs = 377
		goto _test_eof
	_test_eof378:
		cs = 378
		goto _test_eof
	_test_eof379:
		cs = 379
		goto _test_eof
	_test_eof380:
		cs = 380
		goto _test_eof
	_test_eof381:
		cs = 381
		goto _test_eof
	_test_eof382:
		cs = 382
		goto _test_eof
	_test_eof383:
		cs = 383
		goto _test_eof
	_test_eof384:
		cs = 384
		goto _test_eof
	_test_eof385:
		cs = 385
		goto _test_eof
	_test_eof386:
		cs = 386
		goto _test_eof
	_test_eof387:
		cs = 387
		goto _test_eof
	_test_eof388:
		cs = 388
		goto _test_eof
	_test_eof389:
		cs = 389
		goto _test_eof
	_test_eof390:
		cs = 390
		goto _test_eof
	_test_eof391:
		cs = 391
		goto _test_eof
	_test_eof392:
		cs = 392
		goto _test_eof
	_test_eof393:
		cs = 393
		goto _test_eof
	_test_eof394:
		cs = 394
		goto _test_eof
	_test_eof395:
		cs = 395
		goto _test_eof
	_test_eof396:
		cs = 396
		goto _test_eof
	_test_eof397:
		cs = 397
		goto _test_eof
	_test_eof398:
		cs = 398
		goto _test_eof
	_test_eof399:
		cs = 399
		goto _test_eof
	_test_eof400:
		cs = 400
		goto _test_eof
	_test_eof401:
		cs = 401
		goto _test_eof
	_test_eof402:
		cs = 402
		goto _test_eof
	_test_eof403:
		cs = 403
		goto _test_eof
	_test_eof404:
		cs = 404
		goto _test_eof
	_test_eof405:
		cs = 405
		goto _test_eof
	_test_eof406:
		cs = 406
		goto _test_eof
	_test_eof407:
		cs = 407
		goto _test_eof
	_test_eof408:
		cs = 408
		goto _test_eof
	_test_eof409:
		cs = 409
		goto _test_eof
	_test_eof410:
		cs = 410
		goto _test_eof
	_test_eof411:
		cs = 411
		goto _test_eof
	_test_eof412:
		cs = 412
		goto _test_eof
	_test_eof413:
		cs = 413
		goto _test_eof
	_test_eof414:
		cs = 414
		goto _test_eof
	_test_eof415:
		cs = 415
		goto _test_eof
	_test_eof416:
		cs = 416
		goto _test_eof
	_test_eof417:
		cs = 417
		goto _test_eof
	_test_eof418:
		cs = 418
		goto _test_eof
	_test_eof419:
		cs = 419
		goto _test_eof
	_test_eof420:
		cs = 420
		goto _test_eof
	_test_eof421:
		cs = 421
		goto _test_eof
	_test_eof422:
		cs = 422
		goto _test_eof
	_test_eof423:
		cs = 423
		goto _test_eof
	_test_eof424:
		cs = 424
		goto _test_eof
	_test_eof425:
		cs = 425
		goto _test_eof
	_test_eof426:
		cs = 426
		goto _test_eof
	_test_eof427:
		cs = 427
		goto _test_eof
	_test_eof428:
		cs = 428
		goto _test_eof
	_test_eof429:
		cs = 429
		goto _test_eof
	_test_eof430:
		cs = 430
		goto _test_eof
	_test_eof431:
		cs = 431
		goto _test_eof
	_test_eof432:
		cs = 432
		goto _test_eof
	_test_eof433:
		cs = 433
		goto _test_eof
	_test_eof434:
		cs = 434
		goto _test_eof
	_test_eof457:
		cs = 457
		goto _test_eof
	_test_eof435:
		cs = 435
		goto _test_eof
	_test_eof436:
		cs = 436
		goto _test_eof
	_test_eof437:
		cs = 437
		goto _test_eof
	_test_eof438:
		cs = 438
		goto _test_eof
	_test_eof439:
		cs = 439
		goto _test_eof
	_test_eof440:
		cs = 440
		goto _test_eof
	_test_eof441:
		cs = 441
		goto _test_eof
	_test_eof442:
		cs = 442
		goto _test_eof
	_test_eof443:
		cs = 443
		goto _test_eof
	_test_eof444:
		cs = 444
		goto _test_eof
	_test_eof445:
		cs = 445
		goto _test_eof
	_test_eof446:
		cs = 446
		goto _test_eof
	_test_eof447:
		cs = 447
		goto _test_eof
	_test_eof448:
		cs = 448
		goto _test_eof
	_test_eof449:
		cs = 449
		goto _test_eof
	_test_eof450:
		cs = 450
		goto _test_eof
	_test_eof451:
		cs = 451
		goto _test_eof
	_test_eof452:
		cs = 452
		goto _test_eof
	_test_eof453:
		cs = 453
		goto _test_eof
	_test_eof454:
		cs = 454
		goto _test_eof

	_test_eof:
		{
		}
	_out:
		{
		}
	}

	// line 124 "smtp.ragel"

	if cs == smtp_error {
		return data[p:], fmt.Errorf("Invalid character in pos %d: `%c`.", p, data[p])
	}

	// Not yet a full parse, remeber everything from pb to p if we are
	// recording.
	if cs < smtp_first_final {
		if parser.buffer != nil {
			_, err := parser.buffer.Write(data[pb:p])
			if err != nil {
				return data[p:], fmt.Errorf("Could not buffer data: %s", err)
			}
		}
		parser.cs = cs
		return data[p:], Dangling
	}

	// Full parse:
	parser.cs = smtp_en_main
	return data[p:], nil
}
