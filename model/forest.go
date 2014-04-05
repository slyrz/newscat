package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[3] <= -0.2716 {
		if s[4] <= 1.2448 {
			if s[3] <= -2.1045 {
				if s[0] <= 0.0121 {
					if s[3] <= -4.2476 {
						return 0.00571
					} else {
						return 0.04943
					}
				} else {
					return 0.00000
				}
			} else {
				if s[1] <= 6.5000 {
					if s[5] <= 3.4038 {
						return 0.05729
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= -1.5456 {
						return 0.15942
					} else {
						return 0.39583
					}
				}
			}
		} else {
			if s[6] <= -6.5425 {
				if s[7] <= 0.3889 {
					if s[3] <= -1.0730 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					return 1.00000
				}
			} else {
				if s[5] <= 1.0564 {
					if s[0] <= 0.1389 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					return 1.00000
				}
			}
		}
	} else {
		if s[0] <= 0.3203 {
			if s[4] <= 0.1740 {
				if s[1] <= 5.5000 {
					if s[0] <= 0.2088 {
						return 0.36792
					} else {
						return 0.07317
					}
				} else {
					if s[3] <= 2.5040 {
						return 0.85714
					} else {
						return 0.35484
					}
				}
			} else {
				if s[6] <= 0.8346 {
					if s[1] <= 10.5000 {
						return 0.54000
					} else {
						return 0.97222
					}
				} else {
					if s[6] <= 1.7255 {
						return 0.93243
					} else {
						return 0.99272
					}
				}
			}
		} else {
			if s[4] <= -0.7898 {
				if s[5] <= -9.6139 {
					if s[6] <= 0.8985 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= -2.2444 {
						return 0.00000
					} else {
						return 0.05882
					}
				}
			} else {
				if s[5] <= -2.1776 {
					return 0.00000
				} else {
					if s[3] <= 2.7204 {
						return 0.93333
					} else {
						return 0.48148
					}
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[4] <= -0.4148 {
		if s[7] <= 0.4808 {
			if s[4] <= -0.8557 {
				if s[5] <= -4.4342 {
					if s[3] <= -0.9614 {
						return 0.00015
					} else {
						return 0.10870
					}
				} else {
					if s[6] <= 5.7963 {
						return 0.02368
					} else {
						return 0.66667
					}
				}
			} else {
				if s[3] <= -0.9612 {
					if s[5] <= 2.1508 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= -0.5297 {
						return 0.69231
					} else {
						return 0.21053
					}
				}
			}
		} else {
			if s[0] <= 0.5000 {
				if s[7] <= 0.6833 {
					if s[1] <= 6.5000 {
						return 0.40000
					} else {
						return 1.00000
					}
				} else {
					if s[7] <= 0.8516 {
						return 0.88235
					} else {
						return 1.00000
					}
				}
			} else {
				return 0.00000
			}
		}
	} else {
		if s[2] <= 0.5000 {
			if s[5] <= 0.8707 {
				if s[1] <= 4.5000 {
					if s[6] <= -0.4998 {
						return 0.11765
					} else {
						return 0.75610
					}
				} else {
					if s[0] <= 0.1809 {
						return 0.94915
					} else {
						return 0.58621
					}
				}
			} else {
				if s[0] <= 0.4457 {
					if s[6] <= -0.1776 {
						return 0.47059
					} else {
						return 0.97242
					}
				} else {
					if s[4] <= -0.0906 {
						return 0.66667
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[6] <= 1.8217 {
				if s[0] <= 0.2064 {
					if s[3] <= -2.6603 {
						return 0.25000
					} else {
						return 0.95652
					}
				} else {
					if s[5] <= 2.6558 {
						return 0.71429
					} else {
						return 0.29412
					}
				}
			} else {
				if s[6] <= 2.4924 {
					if s[6] <= 2.4886 {
						return 0.99160
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= 2.4923 {
						return 0.98701
					} else {
						return 0.99938
					}
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[3] <= -0.2731 {
		if s[7] <= 0.4226 {
			if s[3] <= -1.0120 {
				if s[5] <= 1.3194 {
					if s[5] <= 0.5840 {
						return 0.00069
					} else {
						return 0.20000
					}
				} else {
					if s[6] <= 0.1832 {
						return 0.11111
					} else {
						return 1.00000
					}
				}
			} else {
				if s[8] <= 0.5000 {
					return 0.00000
				} else {
					if s[4] <= 2.3707 {
						return 0.23232
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[7] <= 0.7386 {
				if s[4] <= -4.8790 {
					return 0.00000
				} else {
					return 1.00000
				}
			} else {
				if s[0] <= 0.5000 {
					if s[3] <= -1.7421 {
						return 0.95122
					} else {
						return 0.36364
					}
				} else {
					return 0.00000
				}
			}
		}
	} else {
		if s[1] <= 8.5000 {
			if s[0] <= 0.2804 {
				if s[0] <= 0.0034 {
					if s[6] <= -0.6916 {
						return 0.28319
					} else {
						return 0.88615
					}
				} else {
					if s[4] <= 0.0054 {
						return 0.32500
					} else {
						return 0.96697
					}
				}
			} else {
				if s[4] <= -0.4757 {
					return 0.00000
				} else {
					if s[6] <= -0.4757 {
						return 0.10000
					} else {
						return 0.83333
					}
				}
			}
		} else {
			if s[3] <= 1.0331 {
				if s[1] <= 77.5000 {
					if s[6] <= 2.6772 {
						return 0.93443
					} else {
						return 0.55556
					}
				} else {
					return 0.00000
				}
			} else {
				if s[6] <= 1.8217 {
					if s[3] <= 9.4562 {
						return 0.94724
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= -1.1959 {
						return 0.14286
					} else {
						return 0.99858
					}
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[3] <= -0.4035 {
		if s[6] <= 3.5171 {
			if s[7] <= 0.5635 {
				if s[6] <= 1.2242 {
					if s[4] <= 3.4225 {
						return 0.00078
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= 1.9923 {
						return 0.60000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 4.5000 {
					return 1.00000
				} else {
					if s[6] <= -9.5202 {
						return 0.62667
					} else {
						return 0.21053
					}
				}
			}
		} else {
			return 1.00000
		}
	} else {
		if s[4] <= 0.1740 {
			if s[4] <= -1.6099 {
				if s[4] <= -2.5009 {
					return 0.00000
				} else {
					if s[6] <= -0.6116 {
						return 0.00000
					} else {
						return 0.66667
					}
				}
			} else {
				if s[7] <= 0.0417 {
					if s[3] <= 3.4138 {
						return 0.58333
					} else {
						return 0.10909
					}
				} else {
					if s[8] <= 0.5000 {
						return 1.00000
					} else {
						return 0.80000
					}
				}
			}
		} else {
			if s[0] <= 0.3106 {
				if s[5] <= -10.4613 {
					return 0.00000
				} else {
					if s[5] <= 0.3013 {
						return 0.92264
					} else {
						return 0.98185
					}
				}
			} else {
				if s[6] <= -0.4061 {
					return 0.00000
				} else {
					if s[6] <= 4.1302 {
						return 0.84000
					} else {
						return 0.14286
					}
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[5] <= 0.3652 {
		if s[4] <= -0.4725 {
			if s[6] <= -0.5808 {
				if s[1] <= 6.5000 {
					if s[1] <= 4.5000 {
						return 0.00012
					} else {
						return 0.00466
					}
				} else {
					if s[4] <= -2.1688 {
						return 0.00561
					} else {
						return 0.30841
					}
				}
			} else {
				if s[0] <= 0.0388 {
					if s[5] <= -5.9371 {
						return 0.66667
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= 2.4251 {
						return 0.17647
					} else {
						return 0.04000
					}
				}
			}
		} else {
			if s[1] <= 4.5000 {
				if s[6] <= -0.3646 {
					if s[7] <= 0.5625 {
						return 0.03947
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.0086 {
						return 0.59091
					} else {
						return 0.88462
					}
				}
			} else {
				if s[3] <= 0.8682 {
					if s[1] <= 23.5000 {
						return 0.91111
					} else {
						return 0.28571
					}
				} else {
					if s[0] <= 0.2429 {
						return 0.98381
					} else {
						return 0.66667
					}
				}
			}
		}
	} else {
		if s[4] <= 0.3649 {
			if s[1] <= 5.5000 {
				if s[4] <= -1.6098 {
					return 0.00000
				} else {
					if s[3] <= 2.3803 {
						return 0.68182
					} else {
						return 0.21311
					}
				}
			} else {
				if s[3] <= 2.8800 {
					if s[1] <= 22.5000 {
						return 0.75862
					} else {
						return 0.31250
					}
				} else {
					if s[6] <= 6.1963 {
						return 0.06061
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[1] <= 10.5000 {
				if s[4] <= 2.6459 {
					if s[6] <= 0.6877 {
						return 0.56000
					} else {
						return 0.92288
					}
				} else {
					if s[4] <= 7.5746 {
						return 0.99312
					} else {
						return 0.66667
					}
				}
			} else {
				if s[1] <= 16.5000 {
					if s[5] <= 8.3007 {
						return 0.99150
					} else {
						return 0.80000
					}
				} else {
					return 1.00000
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[6] <= -0.1794 {
		if s[5] <= 0.4207 {
			if s[7] <= 0.4226 {
				if s[4] <= 0.5737 {
					if s[3] <= -0.9612 {
						return 0.00064
					} else {
						return 0.15049
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.35897
					} else {
						return 0.88679
					}
				}
			} else {
				if s[0] <= 0.5000 {
					if s[7] <= 0.6833 {
						return 0.71429
					} else {
						return 0.96875
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[1] <= 9.5000 {
				if s[3] <= 1.4948 {
					if s[5] <= 1.8106 {
						return 0.15789
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 1.6840 {
						return 1.00000
					} else {
						return 0.39831
					}
				}
			} else {
				if s[4] <= -0.1686 {
					if s[1] <= 37.5000 {
						return 0.90000
					} else {
						return 0.00000
					}
				} else {
					if s[1] <= 13.5000 {
						return 0.95652
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.3153 {
			if s[0] <= 0.2118 {
				if s[4] <= -0.2371 {
					if s[6] <= 2.8454 {
						return 0.23256
					} else {
						return 0.59259
					}
				} else {
					if s[4] <= 2.8491 {
						return 0.92916
					} else {
						return 0.99697
					}
				}
			} else {
				if s[3] <= 0.4711 {
					return 0.00000
				} else {
					if s[3] <= 7.3554 {
						return 0.92199
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[0] <= 0.8128 {
				if s[3] <= 3.7828 {
					if s[4] <= 0.4937 {
						return 0.21429
					} else {
						return 0.68750
					}
				} else {
					if s[6] <= 1.4989 {
						return 0.60000
					} else {
						return 1.00000
					}
				}
			} else {
				return 0.00000
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[4] <= -0.4061 {
		if s[7] <= 0.4808 {
			if s[6] <= -0.5717 {
				if s[3] <= -2.2368 {
					if s[1] <= 18.5000 {
						return 0.00020
					} else {
						return 0.02020
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.01087
					} else {
						return 0.19091
					}
				}
			} else {
				if s[0] <= 0.1966 {
					if s[0] <= 0.1016 {
						return 0.22581
					} else {
						return 0.54545
					}
				} else {
					if s[4] <= -2.3075 {
						return 0.00000
					} else {
						return 0.17647
					}
				}
			}
		} else {
			if s[4] <= -5.7248 {
				return 0.00000
			} else {
				if s[3] <= -3.2988 {
					if s[6] <= -6.0587 {
						return 1.00000
					} else {
						return 0.75000
					}
				} else {
					if s[1] <= 10.5000 {
						return 0.78947
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.6185 {
			if s[3] <= -0.4086 {
				if s[5] <= -9.4776 {
					if s[7] <= 0.3625 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[1] <= 16.5000 {
						return 0.00000
					} else {
						return 0.66667
					}
				}
			} else {
				if s[1] <= 6.5000 {
					if s[6] <= 0.1563 {
						return 0.39655
					} else {
						return 0.94118
					}
				} else {
					if s[3] <= 9.4562 {
						return 0.95607
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[4] <= 0.4454 {
				if s[1] <= 0.5000 {
					return 0.00000
				} else {
					if s[6] <= 5.6754 {
						return 0.75000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[3] <= -0.4482 {
					if s[0] <= 0.2004 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= 2.8488 {
						return 0.94652
					} else {
						return 0.99735
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[4] <= -0.2048 {
		if s[4] <= -1.7600 {
			if s[4] <= -2.2435 {
				if s[4] <= -5.3684 {
					if s[1] <= 15.5000 {
						return 0.00000
					} else {
						return 0.00820
					}
				} else {
					if s[8] <= 0.5000 {
						return 0.06426
					} else {
						return 0.00665
					}
				}
			} else {
				if s[8] <= 0.5000 {
					if s[7] <= 0.0833 {
						return 0.00000
					} else {
						return 0.91667
					}
				} else {
					if s[3] <= 1.3591 {
						return 0.01149
					} else {
						return 0.14286
					}
				}
			}
		} else {
			if s[7] <= 0.1707 {
				if s[4] <= -0.8554 {
					if s[3] <= -1.0725 {
						return 0.01351
					} else {
						return 0.38462
					}
				} else {
					if s[6] <= -0.8930 {
						return 0.17308
					} else {
						return 0.55405
					}
				}
			} else {
				if s[4] <= -1.6473 {
					if s[3] <= -1.7421 {
						return 1.00000
					} else {
						return 0.25000
					}
				} else {
					if s[3] <= -0.4392 {
						return 1.00000
					} else {
						return 0.66667
					}
				}
			}
		}
	} else {
		if s[4] <= 1.3753 {
			if s[1] <= 4.5000 {
				if s[6] <= -0.1921 {
					if s[5] <= 0.2530 {
						return 0.00000
					} else {
						return 0.40000
					}
				} else {
					if s[5] <= -0.8756 {
						return 0.60000
					} else {
						return 0.89565
					}
				}
			} else {
				if s[3] <= 3.7472 {
					if s[2] <= 1.5000 {
						return 0.89524
					} else {
						return 0.66667
					}
				} else {
					if s[8] <= 0.5000 {
						return 1.00000
					} else {
						return 0.21429
					}
				}
			}
		} else {
			if s[3] <= -2.9204 {
				if s[7] <= 0.3889 {
					return 0.00000
				} else {
					return 1.00000
				}
			} else {
				if s[5] <= -10.4613 {
					return 0.00000
				} else {
					if s[2] <= 0.5000 {
						return 0.94722
					} else {
						return 0.99434
					}
				}
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[6] <= 0.0204 {
		if s[7] <= 0.4202 {
			if s[5] <= 0.4354 {
				if s[5] <= -1.1539 {
					if s[4] <= 0.9511 {
						return 0.00206
					} else {
						return 0.71642
					}
				} else {
					if s[4] <= -0.7898 {
						return 0.01020
					} else {
						return 0.52941
					}
				}
			} else {
				if s[4] <= 0.1323 {
					if s[4] <= -1.7453 {
						return 0.00000
					} else {
						return 0.29091
					}
				} else {
					if s[0] <= 0.3066 {
						return 0.88339
					} else {
						return 0.11111
					}
				}
			}
		} else {
			if s[0] <= 0.5000 {
				if s[1] <= 8.5000 {
					return 1.00000
				} else {
					if s[3] <= -3.1157 {
						return 1.00000
					} else {
						return 0.76190
					}
				}
			} else {
				return 0.00000
			}
		}
	} else {
		if s[4] <= -0.1993 {
			if s[4] <= -0.8581 {
				if s[3] <= 3.2651 {
					if s[4] <= -2.4634 {
						return 0.02439
					} else {
						return 0.27273
					}
				} else {
					return 0.00000
				}
			} else {
				if s[2] <= 1.5000 {
					if s[0] <= 0.2639 {
						return 0.44828
					} else {
						return 0.78571
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[0] <= 0.8128 {
				if s[6] <= 1.8310 {
					if s[5] <= -10.2741 {
						return 0.00000
					} else {
						return 0.89362
					}
				} else {
					if s[1] <= 1.5000 {
						return 0.94224
					} else {
						return 0.99460
					}
				}
			} else {
				return 0.00000
			}
		}
	}
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[3] <= -0.2731 {
		if s[7] <= 0.5635 {
			if s[5] <= 1.2299 {
				if s[4] <= -0.7024 {
					if s[6] <= -1.7703 {
						return 0.00040
					} else {
						return 0.05714
					}
				} else {
					if s[6] <= 0.9719 {
						return 0.12644
					} else {
						return 0.63636
					}
				}
			} else {
				if s[6] <= 1.4302 {
					if s[6] <= 0.8863 {
						return 0.46154
					} else {
						return 0.00000
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[8] <= 0.5000 {
				if s[4] <= -5.4918 {
					return 0.00000
				} else {
					if s[1] <= 8.5000 {
						return 0.94444
					} else {
						return 0.68000
					}
				}
			} else {
				return 0.00000
			}
		}
	} else {
		if s[6] <= 0.1319 {
			if s[1] <= 8.5000 {
				if s[4] <= 1.3929 {
					if s[6] <= -0.8554 {
						return 0.09259
					} else {
						return 0.41935
					}
				} else {
					if s[5] <= -2.1501 {
						return 0.85714
					} else {
						return 0.53846
					}
				}
			} else {
				if s[1] <= 10.5000 {
					if s[8] <= 0.5000 {
						return 1.00000
					} else {
						return 0.60526
					}
				} else {
					if s[0] <= 0.2088 {
						return 0.98013
					} else {
						return 0.33333
					}
				}
			}
		} else {
			if s[4] <= -0.5005 {
				if s[3] <= 2.9145 {
					if s[0] <= 0.8115 {
						return 0.41379
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= -9.5067 {
						return 0.14286
					} else {
						return 0.02222
					}
				}
			} else {
				if s[0] <= 0.3153 {
					if s[1] <= 4.5000 {
						return 0.95369
					} else {
						return 0.99569
					}
				} else {
					if s[5] <= 1.1239 {
						return 0.95238
					} else {
						return 0.44000
					}
				}
			}
		}
	}
}
