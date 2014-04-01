package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[5] <= 0.1120 {
		if s[6] <= 0.2037 {
			if s[3] <= -1.1055 {
				if s[3] <= -2.7147 {
					if s[1] <= 5.5000 {
						return 0.00020
					} else {
						return 0.00748
					}
				} else {
					if s[1] <= 5.5000 {
						return 0.01802
					} else {
						return 0.19802
					}
				}
			} else {
				if s[4] <= -0.1529 {
					if s[6] <= -1.6584 {
						return 0.11240
					} else {
						return 0.39024
					}
				} else {
					if s[0] <= 0.1671 {
						return 0.57500
					} else {
						return 0.09375
					}
				}
			}
		} else {
			if s[0] <= 0.3144 {
				if s[6] <= 1.7729 {
					if s[5] <= -0.2880 {
						return 0.55357
					} else {
						return 0.90476
					}
				} else {
					if s[1] <= 19.5000 {
						return 0.83962
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= -0.1799 {
					if s[4] <= -1.4810 {
						return 0.00000
					} else {
						return 0.22222
					}
				} else {
					if s[5] <= -5.1705 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.3058 {
			if s[4] <= 0.1336 {
				if s[4] <= -1.4224 {
					if s[5] <= 0.2375 {
						return 1.00000
					} else {
						return 0.06667
					}
				} else {
					if s[2] <= 1.5000 {
						return 0.48120
					} else {
						return 0.22222
					}
				}
			} else {
				if s[4] <= 1.8119 {
					if s[2] <= 0.5000 {
						return 0.83462
					} else {
						return 0.92795
					}
				} else {
					if s[1] <= 12.5000 {
						return 0.95332
					} else {
						return 0.99779
					}
				}
			}
		} else {
			if s[1] <= 2.5000 {
				if s[5] <= 0.1207 {
					return 1.00000
				} else {
					if s[0] <= 0.3376 {
						return 0.25000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[3] <= 3.0232 {
					if s[4] <= -0.0842 {
						return 0.07143
					} else {
						return 0.68000
					}
				} else {
					if s[0] <= 0.6635 {
						return 0.87500
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[3] <= -0.2376 {
		if s[5] <= -0.8728 {
			if s[6] <= 0.6894 {
				if s[6] <= -1.4804 {
					if s[1] <= 5.5000 {
						return 0.00043
					} else {
						return 0.02112
					}
				} else {
					if s[3] <= -1.0105 {
						return 0.06630
					} else {
						return 0.82353
					}
				}
			} else {
				if s[1] <= 12.0000 {
					if s[5] <= -3.4842 {
						return 0.11111
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= 1.1326 {
						return 0.50000
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[1] <= 9.5000 {
				if s[1] <= 6.5000 {
					if s[3] <= -1.8813 {
						return 0.00000
					} else {
						return 0.17526
					}
				} else {
					if s[0] <= 0.0172 {
						return 0.15000
					} else {
						return 0.72727
					}
				}
			} else {
				if s[5] <= 0.2440 {
					if s[2] <= 1.5000 {
						return 0.71429
					} else {
						return 0.32000
					}
				} else {
					if s[0] <= 0.0731 {
						return 0.95312
					} else {
						return 0.27273
					}
				}
			}
		}
	} else {
		if s[0] <= 0.2829 {
			if s[6] <= 0.7388 {
				if s[2] <= 0.5000 {
					if s[6] <= -1.0523 {
						return 0.36747
					} else {
						return 0.59375
					}
				} else {
					if s[0] <= 0.2097 {
						return 0.88866
					} else {
						return 0.00000
					}
				}
			} else {
				if s[4] <= -0.1747 {
					if s[0] <= 0.2375 {
						return 0.38889
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= 1.8696 {
						return 0.93478
					} else {
						return 0.99101
					}
				}
			}
		} else {
			if s[0] <= 0.5708 {
				if s[4] <= 0.2342 {
					if s[3] <= 0.6424 {
						return 0.00000
					} else {
						return 0.15909
					}
				} else {
					if s[3] <= 3.0194 {
						return 0.48980
					} else {
						return 0.94118
					}
				}
			} else {
				if s[0] <= 0.8655 {
					if s[3] <= 2.6336 {
						return 0.14894
					} else {
						return 0.00000
					}
				} else {
					return 0.00000
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[4] <= 0.1248 {
		if s[1] <= 6.5000 {
			if s[3] <= -1.1297 {
				if s[6] <= -1.6415 {
					if s[2] <= 1.5000 {
						return 0.00036
					} else {
						return 0.03509
					}
				} else {
					if s[4] <= -1.2335 {
						return 0.00909
					} else {
						return 0.53846
					}
				}
			} else {
				if s[1] <= 1.5000 {
					if s[3] <= -1.1001 {
						return 1.00000
					} else {
						return 0.08889
					}
				} else {
					if s[1] <= 4.5000 {
						return 0.25333
					} else {
						return 0.15385
					}
				}
			}
		} else {
			if s[5] <= -1.6709 {
				if s[0] <= 0.0731 {
					if s[0] <= 0.0728 {
						return 0.03830
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= -1.9472 {
						return 0.00000
					} else {
						return 0.16667
					}
				}
			} else {
				if s[4] <= -0.3737 {
					if s[2] <= 0.5000 {
						return 0.40476
					} else {
						return 0.15556
					}
				} else {
					if s[3] <= 2.1196 {
						return 0.68182
					} else {
						return 0.11111
					}
				}
			}
		}
	} else {
		if s[4] <= 1.4087 {
			if s[1] <= 9.5000 {
				if s[3] <= -0.0106 {
					if s[1] <= 5.5000 {
						return 0.13846
					} else {
						return 0.55556
					}
				} else {
					if s[3] <= 2.6914 {
						return 0.75145
					} else {
						return 0.39344
					}
				}
			} else {
				if s[2] <= 2.5000 {
					if s[0] <= 0.1315 {
						return 0.93204
					} else {
						return 0.79167
					}
				} else {
					if s[6] <= -8.3615 {
						return 0.07692
					} else {
						return 0.80000
					}
				}
			}
		} else {
			if s[1] <= 13.5000 {
				if s[6] <= 1.0646 {
					if s[4] <= 2.9937 {
						return 0.56154
					} else {
						return 0.91667
					}
				} else {
					if s[3] <= 0.9776 {
						return 0.84483
					} else {
						return 0.98035
					}
				}
			} else {
				if s[6] <= 1.8961 {
					if s[3] <= -2.0650 {
						return 0.00000
					} else {
						return 0.96524
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.99138
					} else {
						return 1.00000
					}
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[3] <= -0.2376 {
		if s[6] <= -1.4804 {
			if s[3] <= -2.4556 {
				if s[3] <= -3.3949 {
					if s[3] <= -4.0028 {
						return 0.00058
					} else {
						return 0.01495
					}
				} else {
					if s[3] <= -3.3933 {
						return 1.00000
					} else {
						return 0.05316
					}
				}
			} else {
				if s[4] <= 0.0652 {
					if s[1] <= 4.5000 {
						return 0.02929
					} else {
						return 0.23443
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.34545
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[4] <= 0.2563 {
				if s[4] <= -1.4662 {
					if s[6] <= -1.4074 {
						return 0.22222
					} else {
						return 0.01325
					}
				} else {
					if s[5] <= 0.1580 {
						return 0.17647
					} else {
						return 0.61905
					}
				}
			} else {
				if s[1] <= 0.5000 {
					return 0.00000
				} else {
					if s[0] <= 0.1647 {
						return 0.87324
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.1160 {
			if s[4] <= 1.1606 {
				if s[1] <= 6.5000 {
					if s[1] <= 1.5000 {
						return 0.02857
					} else {
						return 0.15672
					}
				} else {
					if s[3] <= 2.6690 {
						return 0.56962
					} else {
						return 0.09091
					}
				}
			} else {
				if s[1] <= 11.5000 {
					if s[0] <= 0.2259 {
						return 0.62832
					} else {
						return 0.08333
					}
				} else {
					if s[3] <= 1.7490 {
						return 0.83929
					} else {
						return 0.97817
					}
				}
			}
		} else {
			if s[4] <= 0.1578 {
				if s[5] <= -0.5374 {
					if s[1] <= 2.5000 {
						return 0.00000
					} else {
						return 0.23810
					}
				} else {
					if s[5] <= 0.7931 {
						return 0.60526
					} else {
						return 0.30263
					}
				}
			} else {
				if s[4] <= 1.0810 {
					if s[0] <= 0.1662 {
						return 0.91304
					} else {
						return 0.65714
					}
				} else {
					if s[0] <= 0.2747 {
						return 0.98887
					} else {
						return 0.89041
					}
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[6] <= -0.4543 {
		if s[3] <= -0.9121 {
			if s[5] <= -1.5396 {
				if s[4] <= -2.3628 {
					if s[5] <= -2.1966 {
						return 0.00056
					} else {
						return 0.04040
					}
				} else {
					if s[4] <= -1.1797 {
						return 0.09145
					} else {
						return 0.20809
					}
				}
			} else {
				if s[5] <= -1.5323 {
					return 1.00000
				} else {
					if s[2] <= 4.5000 {
						return 0.12308
					} else {
						return 0.44444
					}
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[1] <= 4.5000 {
					if s[5] <= 1.1254 {
						return 0.08213
					} else {
						return 0.22535
					}
				} else {
					if s[4] <= -0.8147 {
						return 0.02632
					} else {
						return 0.76531
					}
				}
			} else {
				if s[0] <= 0.2066 {
					if s[6] <= -9.7910 {
						return 0.77713
					} else {
						return 0.91262
					}
				} else {
					if s[1] <= 0.5000 {
						return 0.30000
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.3237 {
			if s[3] <= -1.1599 {
				if s[5] <= 0.0845 {
					if s[1] <= 1.5000 {
						return 0.25000
					} else {
						return 0.05172
					}
				} else {
					if s[0] <= 0.1317 {
						return 0.95238
					} else {
						return 0.00000
					}
				}
			} else {
				if s[5] <= 1.3461 {
					if s[0] <= 0.2377 {
						return 0.87805
					} else {
						return 0.47222
					}
				} else {
					if s[6] <= 1.9967 {
						return 0.93548
					} else {
						return 0.99429
					}
				}
			}
		} else {
			if s[0] <= 0.5708 {
				if s[4] <= 0.1503 {
					if s[5] <= -0.2571 {
						return 0.25000
					} else {
						return 0.04762
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.48387
					} else {
						return 1.00000
					}
				}
			} else {
				if s[1] <= 11.5000 {
					if s[2] <= 0.5000 {
						return 0.01667
					} else {
						return 0.20000
					}
				} else {
					return 1.00000
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[5] <= -0.0115 {
		if s[3] <= -0.9201 {
			if s[5] <= -1.8380 {
				if s[3] <= -2.7871 {
					if s[0] <= 0.0609 {
						return 0.00527
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.0637 {
						return 0.14684
					} else {
						return 0.01333
					}
				}
			} else {
				if s[6] <= -1.0018 {
					if s[1] <= 6.5000 {
						return 0.04575
					} else {
						return 0.17978
					}
				} else {
					if s[3] <= -1.1851 {
						return 0.24324
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[6] <= 0.8668 {
				if s[2] <= 0.5000 {
					if s[3] <= 3.3696 {
						return 0.21474
					} else {
						return 0.66667
					}
				} else {
					if s[2] <= 2.5000 {
						return 0.54918
					} else {
						return 0.33333
					}
				}
			} else {
				if s[0] <= 0.2834 {
					if s[1] <= 13.5000 {
						return 0.74766
					} else {
						return 0.96020
					}
				} else {
					return 0.00000
				}
			}
		}
	} else {
		if s[0] <= 0.3106 {
			if s[4] <= 0.1210 {
				if s[0] <= 0.1688 {
					if s[6] <= -1.0989 {
						return 0.14706
					} else {
						return 0.64835
					}
				} else {
					if s[0] <= 0.2711 {
						return 0.09615
					} else {
						return 0.44444
					}
				}
			} else {
				if s[6] <= 1.4297 {
					if s[4] <= 3.2806 {
						return 0.84557
					} else {
						return 0.96891
					}
				} else {
					if s[4] <= 1.0434 {
						return 0.88406
					} else {
						return 0.98975
					}
				}
			}
		} else {
			if s[4] <= -0.8714 {
				if s[6] <= 3.7679 {
					return 0.00000
				} else {
					if s[4] <= -2.5520 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[1] <= 10.5000 {
					if s[0] <= 0.6992 {
						return 0.55102
					} else {
						return 0.00000
					}
				} else {
					return 1.00000
				}
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[4] <= 0.0771 {
		if s[6] <= -1.1700 {
			if s[4] <= -2.0175 {
				if s[1] <= 11.5000 {
					if s[5] <= -3.2221 {
						return 0.00033
					} else {
						return 0.02110
					}
				} else {
					if s[0] <= 0.0381 {
						return 0.03125
					} else {
						return 0.00299
					}
				}
			} else {
				if s[3] <= -6.0051 {
					if s[1] <= 4.0000 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.1559 {
						return 0.20131
					} else {
						return 0.02586
					}
				}
			}
		} else {
			if s[3] <= -1.1055 {
				if s[6] <= -1.1586 {
					return 1.00000
				} else {
					if s[3] <= -1.9319 {
						return 0.00000
					} else {
						return 0.11290
					}
				}
			} else {
				if s[0] <= 0.4090 {
					if s[3] <= 2.6653 {
						return 0.56442
					} else {
						return 0.16667
					}
				} else {
					if s[1] <= 11.0000 {
						return 0.05479
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.2111 {
			if s[6] <= -0.2119 {
				if s[5] <= -0.3527 {
					if s[4] <= 2.8546 {
						return 0.48387
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= 2.1014 {
						return 0.78400
					} else {
						return 0.91275
					}
				}
			} else {
				if s[3] <= -2.0650 {
					return 0.00000
				} else {
					if s[5] <= 0.8362 {
						return 0.92086
					} else {
						return 0.99153
					}
				}
			}
		} else {
			if s[3] <= -0.0519 {
				return 0.00000
			} else {
				if s[0] <= 0.4073 {
					if s[5] <= 0.9940 {
						return 0.52941
					} else {
						return 0.88276
					}
				} else {
					if s[2] <= 4.5000 {
						return 0.25714
					} else {
						return 1.00000
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[3] <= -0.5357 {
		if s[4] <= -1.1248 {
			if s[5] <= -0.9279 {
				if s[4] <= -2.7140 {
					if s[0] <= 0.0040 {
						return 0.00306
					} else {
						return 0.00004
					}
				} else {
					if s[1] <= 4.5000 {
						return 0.02083
					} else {
						return 0.15894
					}
				}
			} else {
				if s[0] <= 0.3492 {
					if s[1] <= 6.5000 {
						return 0.04412
					} else {
						return 0.30769
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[0] <= 0.1160 {
				if s[1] <= 4.5000 {
					if s[0] <= 0.1053 {
						return 0.05983
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= -0.1579 {
						return 0.52239
					} else {
						return 0.82857
					}
				}
			} else {
				if s[1] <= 22.5000 {
					return 0.00000
				} else {
					if s[2] <= 3.0000 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.1209 {
			if s[2] <= 0.5000 {
				if s[1] <= 4.5000 {
					if s[3] <= 5.5115 {
						return 0.11439
					} else {
						return 0.75000
					}
				} else {
					if s[0] <= 0.4148 {
						return 0.61224
					} else {
						return 0.02857
					}
				}
			} else {
				if s[4] <= 1.3101 {
					if s[5] <= 2.0479 {
						return 0.47134
					} else {
						return 0.08000
					}
				} else {
					if s[1] <= 7.5000 {
						return 0.68966
					} else {
						return 0.96512
					}
				}
			}
		} else {
			if s[4] <= 0.1564 {
				if s[5] <= 1.2719 {
					if s[4] <= -0.5535 {
						return 0.13725
					} else {
						return 0.51613
					}
				} else {
					if s[4] <= -1.4161 {
						return 0.11765
					} else {
						return 0.71698
					}
				}
			} else {
				if s[6] <= 1.6767 {
					if s[1] <= 4.5000 {
						return 0.76259
					} else {
						return 0.95355
					}
				} else {
					if s[4] <= 1.5600 {
						return 0.94709
					} else {
						return 0.99102
					}
				}
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[5] <= -0.0197 {
		if s[4] <= -0.0276 {
			if s[4] <= -1.6908 {
				if s[4] <= -2.7798 {
					if s[5] <= -2.1952 {
						return 0.00045
					} else {
						return 0.07282
					}
				} else {
					if s[4] <= -2.7765 {
						return 1.00000
					} else {
						return 0.05993
					}
				}
			} else {
				if s[5] <= -5.7836 {
					if s[4] <= -1.6775 {
						return 1.00000
					} else {
						return 0.28436
					}
				} else {
					if s[5] <= -1.5327 {
						return 0.06818
					} else {
						return 0.25185
					}
				}
			}
		} else {
			if s[1] <= 4.5000 {
				if s[4] <= 2.6821 {
					if s[5] <= -0.9357 {
						return 0.20438
					} else {
						return 0.43333
					}
				} else {
					if s[3] <= 3.0223 {
						return 0.83333
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= 1.7447 {
					if s[2] <= 4.0000 {
						return 0.75155
					} else {
						return 0.22222
					}
				} else {
					if s[6] <= -3.1077 {
						return 0.88235
					} else {
						return 0.99020
					}
				}
			}
		}
	} else {
		if s[5] <= 1.4367 {
			if s[1] <= 5.5000 {
				if s[4] <= 0.0888 {
					if s[4] <= -1.4642 {
						return 0.00000
					} else {
						return 0.31818
					}
				} else {
					if s[4] <= 2.1008 {
						return 0.87234
					} else {
						return 0.69444
					}
				}
			} else {
				if s[4] <= 0.1194 {
					if s[2] <= 1.5000 {
						return 0.61111
					} else {
						return 0.18750
					}
				} else {
					if s[1] <= 62.0000 {
						return 0.95019
					} else {
						return 0.55556
					}
				}
			}
		} else {
			if s[4] <= 0.1282 {
				if s[5] <= 2.0592 {
					if s[1] <= 11.5000 {
						return 0.32353
					} else {
						return 0.81818
					}
				} else {
					if s[4] <= -1.1302 {
						return 0.01754
					} else {
						return 0.32787
					}
				}
			} else {
				if s[5] <= 1.9741 {
					if s[3] <= 4.7798 {
						return 0.91875
					} else {
						return 0.00000
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.95538
					} else {
						return 0.98972
					}
				}
			}
		}
	}
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[3] <= -0.2376 {
		if s[3] <= -2.3161 {
			if s[1] <= 5.5000 {
				if s[1] <= 4.5000 {
					if s[2] <= 0.5000 {
						return 0.00011
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.0358 {
						return 0.01732
					} else {
						return 0.00000
					}
				}
			} else {
				if s[0] <= 0.0040 {
					if s[6] <= -11.1972 {
						return 0.09677
					} else {
						return 0.02517
					}
				} else {
					if s[5] <= -1.9145 {
						return 0.00000
					} else {
						return 0.12000
					}
				}
			}
		} else {
			if s[5] <= 0.2440 {
				if s[6] <= 0.0636 {
					if s[4] <= 1.9011 {
						return 0.15889
					} else {
						return 0.80000
					}
				} else {
					if s[1] <= 12.0000 {
						return 0.17857
					} else {
						return 0.67742
					}
				}
			} else {
				if s[0] <= 0.2030 {
					if s[1] <= 7.5000 {
						return 0.40000
					} else {
						return 0.86885
					}
				} else {
					if s[1] <= 2.5000 {
						return 0.00000
					} else {
						return 0.11111
					}
				}
			}
		}
	} else {
		if s[6] <= -0.4584 {
			if s[1] <= 10.5000 {
				if s[0] <= 0.2066 {
					if s[4] <= 1.2226 {
						return 0.13043
					} else {
						return 0.72549
					}
				} else {
					if s[4] <= -0.0332 {
						return 0.00000
					} else {
						return 0.24242
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[4] <= 0.2552 {
						return 0.00000
					} else {
						return 0.72727
					}
				} else {
					if s[3] <= 1.5949 {
						return 0.67593
					} else {
						return 0.92857
					}
				}
			}
		} else {
			if s[4] <= 0.1072 {
				if s[3] <= 2.6476 {
					if s[4] <= -1.3739 {
						return 0.07692
					} else {
						return 0.57407
					}
				} else {
					if s[3] <= 3.3679 {
						return 0.00000
					} else {
						return 0.37500
					}
				}
			} else {
				if s[4] <= 1.0810 {
					if s[0] <= 0.3974 {
						return 0.89418
					} else {
						return 0.46154
					}
				} else {
					if s[3] <= 0.9776 {
						return 0.92105
					} else {
						return 0.98886
					}
				}
			}
		}
	}
}
