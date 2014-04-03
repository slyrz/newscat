package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[5] <= -0.2474 {
		if s[3] <= -1.2479 {
			if s[3] <= -1.9838 {
				if s[3] <= -3.1580 {
					if s[4] <= -0.4235 {
						return 0.00054
					} else {
						return 0.61290
					}
				} else {
					if s[3] <= -3.1458 {
						return 1.00000
					} else {
						return 0.05212
					}
				}
			} else {
				if s[1] <= 5.5000 {
					if s[4] <= -1.6455 {
						return 0.03226
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= -1.8195 {
						return 0.23077
					} else {
						return 0.69231
					}
				}
			}
		} else {
			if s[4] <= 1.2626 {
				if s[3] <= 0.9862 {
					if s[2] <= 0.5000 {
						return 0.30275
					} else {
						return 0.50000
					}
				} else {
					if s[4] <= -1.1770 {
						return 0.01198
					} else {
						return 0.20968
					}
				}
			} else {
				if s[1] <= 7.5000 {
					if s[4] <= 3.1852 {
						return 0.71875
					} else {
						return 0.95000
					}
				} else {
					if s[1] <= 13.5000 {
						return 0.93478
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[3] <= -0.8137 {
			if s[1] <= 17.0000 {
				return 0.00000
			} else {
				if s[5] <= 2.2915 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		} else {
			if s[4] <= -0.4509 {
				if s[2] <= 0.5000 {
					if s[6] <= 1.2634 {
						return 0.03158
					} else {
						return 0.31111
					}
				} else {
					if s[4] <= -0.6882 {
						return 0.19231
					} else {
						return 1.00000
					}
				}
			} else {
				if s[0] <= 0.3058 {
					if s[0] <= 0.2111 {
						return 0.97685
					} else {
						return 0.87425
					}
				} else {
					if s[6] <= -0.3621 {
						return 0.21053
					} else {
						return 0.82609
					}
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[3] <= -0.1801 {
		if s[6] <= 0.7958 {
			if s[4] <= -1.3195 {
				if s[4] <= -3.0444 {
					if s[3] <= -2.8652 {
						return 0.00034
					} else {
						return 0.03333
					}
				} else {
					if s[0] <= 0.0116 {
						return 0.10248
					} else {
						return 0.00565
					}
				}
			} else {
				if s[0] <= 0.1149 {
					if s[5] <= -0.0652 {
						return 0.32121
					} else {
						return 0.83333
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[3] <= -1.2193 {
				return 0.00000
			} else {
				if s[5] <= 2.1952 {
					if s[4] <= 0.7982 {
						return 0.83333
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.2869 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[2] <= 0.5000 {
			if s[5] <= -0.1753 {
				if s[0] <= 0.2482 {
					if s[1] <= 2.5000 {
						return 0.32759
					} else {
						return 0.67677
					}
				} else {
					if s[1] <= 11.5000 {
						return 0.04706
					} else {
						return 1.00000
					}
				}
			} else {
				if s[5] <= 1.2870 {
					if s[4] <= -0.1227 {
						return 0.13043
					} else {
						return 0.92105
					}
				} else {
					if s[4] <= -0.0474 {
						return 0.21296
					} else {
						return 0.95032
					}
				}
			}
		} else {
			if s[6] <= -0.6579 {
				if s[4] <= 1.7487 {
					if s[0] <= 0.1640 {
						return 0.58571
					} else {
						return 0.11429
					}
				} else {
					if s[3] <= 7.1726 {
						return 0.99180
					} else {
						return 0.84211
					}
				}
			} else {
				if s[0] <= 0.7851 {
					if s[5] <= 1.6982 {
						return 0.94529
					} else {
						return 0.99636
					}
				} else {
					return 0.00000
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[6] <= -0.5053 {
		if s[5] <= 0.9682 {
			if s[0] <= 0.2246 {
				if s[2] <= 0.5000 {
					if s[4] <= 0.6001 {
						return 0.01292
					} else {
						return 0.70000
					}
				} else {
					if s[5] <= -0.1030 {
						return 0.06890
					} else {
						return 0.55814
					}
				}
			} else {
				if s[6] <= -1.4600 {
					if s[4] <= -2.4820 {
						return 0.00000
					} else {
						return 0.06863
					}
				} else {
					if s[0] <= 0.6536 {
						return 0.04545
					} else {
						return 0.33333
					}
				}
			}
		} else {
			if s[1] <= 9.5000 {
				if s[2] <= 0.5000 {
					if s[4] <= 0.6656 {
						return 0.01282
					} else {
						return 0.59574
					}
				} else {
					if s[0] <= 0.3056 {
						return 0.60000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[3] <= 5.6217 {
					if s[4] <= -0.7951 {
						return 0.00000
					} else {
						return 0.93237
					}
				} else {
					if s[0] <= 0.0692 {
						return 0.98684
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 1.5265 {
			if s[3] <= -1.2193 {
				if s[4] <= -0.4188 {
					return 0.00000
				} else {
					if s[1] <= 2.5000 {
						return 0.50000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[4] <= -1.2827 {
					if s[2] <= 0.5000 {
						return 0.03448
					} else {
						return 0.33333
					}
				} else {
					if s[3] <= 3.3986 {
						return 0.95475
					} else {
						return 0.74561
					}
				}
			}
		} else {
			if s[4] <= -0.0474 {
				if s[0] <= 0.5708 {
					if s[6] <= 3.7105 {
						return 0.55263
					} else {
						return 0.29167
					}
				} else {
					if s[6] <= 5.7634 {
						return 0.00000
					} else {
						return 0.10000
					}
				}
			} else {
				if s[3] <= -0.9473 {
					if s[5] <= -2.0115 {
						return 0.00000
					} else {
						return 0.50000
					}
				} else {
					if s[4] <= 2.1309 {
						return 0.90149
					} else {
						return 0.99331
					}
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[3] <= -0.7306 {
		if s[3] <= -2.1790 {
			if s[6] <= 0.6102 {
				if s[4] <= 0.1136 {
					if s[0] <= 0.0064 {
						return 0.00614
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= -7.0327 {
						return 0.68750
					} else {
						return 0.07143
					}
				}
			} else {
				if s[6] <= 0.6648 {
					return 1.00000
				} else {
					return 0.00000
				}
			}
		} else {
			if s[1] <= 8.5000 {
				if s[3] <= -0.7908 {
					if s[1] <= 5.5000 {
						return 0.01190
					} else {
						return 0.17241
					}
				} else {
					if s[6] <= -3.3833 {
						return 0.10000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= -1.7908 {
					if s[5] <= -1.8802 {
						return 0.72340
					} else {
						return 0.09524
					}
				} else {
					if s[5] <= 1.5649 {
						return 0.00000
					} else {
						return 0.75000
					}
				}
			}
		}
	} else {
		if s[5] <= 0.0454 {
			if s[1] <= 5.5000 {
				if s[4] <= 0.3576 {
					if s[4] <= -1.2057 {
						return 0.00592
					} else {
						return 0.25000
					}
				} else {
					if s[6] <= 1.1443 {
						return 0.23684
					} else {
						return 0.90698
					}
				}
			} else {
				if s[6] <= -2.8200 {
					if s[0] <= 0.1815 {
						return 0.73729
					} else {
						return 0.05000
					}
				} else {
					if s[4] <= -0.7674 {
						return 0.43750
					} else {
						return 0.98182
					}
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[4] <= -0.1224 {
					if s[6] <= -0.5461 {
						return 0.04950
					} else {
						return 0.24359
					}
				} else {
					if s[3] <= 2.5542 {
						return 0.87963
					} else {
						return 0.95177
					}
				}
			} else {
				if s[4] <= -0.6781 {
					if s[3] <= 0.9578 {
						return 0.66667
					} else {
						return 0.06250
					}
				} else {
					if s[6] <= -8.6668 {
						return 0.90043
					} else {
						return 0.99806
					}
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[4] <= -0.4352 {
		if s[5] <= 1.0694 {
			if s[4] <= -2.1882 {
				if s[3] <= -5.0175 {
					if s[0] <= 0.0064 {
						return 0.00228
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= -4.9972 {
						return 0.37500
					} else {
						return 0.00728
					}
				}
			} else {
				if s[4] <= -2.1877 {
					return 1.00000
				} else {
					if s[3] <= -2.0003 {
						return 0.07527
					} else {
						return 0.18367
					}
				}
			}
		} else {
			if s[0] <= 0.3884 {
				if s[3] <= 1.0697 {
					if s[3] <= -1.0131 {
						return 0.00000
					} else {
						return 0.70000
					}
				} else {
					if s[1] <= 4.5000 {
						return 0.33333
					} else {
						return 0.03846
					}
				}
			} else {
				if s[4] <= -2.0276 {
					return 0.00000
				} else {
					if s[1] <= 5.0000 {
						return 0.00000
					} else {
						return 0.50000
					}
				}
			}
		}
	} else {
		if s[3] <= -0.1485 {
			if s[5] <= -4.6889 {
				if s[1] <= 5.5000 {
					if s[6] <= -9.8978 {
						return 0.20000
					} else {
						return 0.88889
					}
				} else {
					if s[6] <= -9.0659 {
						return 0.90909
					} else {
						return 1.00000
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[1] <= 8.0000 {
						return 0.12500
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= -6.8437 {
						return 0.33333
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[6] <= -0.1334 {
				if s[0] <= 0.2280 {
					if s[4] <= 2.5332 {
						return 0.67832
					} else {
						return 0.95222
					}
				} else {
					if s[3] <= 2.2252 {
						return 1.00000
					} else {
						return 0.16667
					}
				}
			} else {
				if s[4] <= 0.8420 {
					if s[5] <= -1.0942 {
						return 0.52632
					} else {
						return 0.85385
					}
				} else {
					if s[4] <= 3.1127 {
						return 0.96382
					} else {
						return 0.99748
					}
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[6] <= -0.5088 {
		if s[4] <= 0.4965 {
			if s[1] <= 6.5000 {
				if s[3] <= -1.8144 {
					if s[5] <= -2.3696 {
						return 0.00034
					} else {
						return 0.03333
					}
				} else {
					if s[4] <= -2.0053 {
						return 0.00000
					} else {
						return 0.08434
					}
				}
			} else {
				if s[4] <= -1.5081 {
					if s[3] <= -4.5584 {
						return 0.00220
					} else {
						return 0.07458
					}
				} else {
					if s[1] <= 22.5000 {
						return 0.62222
					} else {
						return 0.05882
					}
				}
			}
		} else {
			if s[4] <= 2.0439 {
				if s[3] <= -1.2193 {
					if s[1] <= 5.0000 {
						return 0.00000
					} else {
						return 0.50000
					}
				} else {
					if s[6] <= -2.8644 {
						return 0.62162
					} else {
						return 0.27586
					}
				}
			} else {
				if s[1] <= 7.5000 {
					if s[3] <= 4.8656 {
						return 0.76190
					} else {
						return 0.31579
					}
				} else {
					if s[1] <= 15.5000 {
						return 0.87097
					} else {
						return 0.98942
					}
				}
			}
		}
	} else {
		if s[2] <= 0.5000 {
			if s[4] <= -0.4824 {
				if s[0] <= 0.1710 {
					if s[6] <= -0.0176 {
						return 1.00000
					} else {
						return 0.28571
					}
				} else {
					if s[0] <= 0.8681 {
						return 0.11429
					} else {
						return 0.00000
					}
				}
			} else {
				if s[5] <= 1.7588 {
					if s[1] <= 6.5000 {
						return 0.77660
					} else {
						return 0.97403
					}
				} else {
					if s[6] <= -0.1541 {
						return 0.20000
					} else {
						return 0.97172
					}
				}
			}
		} else {
			if s[5] <= -1.7925 {
				if s[3] <= -0.8962 {
					return 0.00000
				} else {
					if s[5] <= -1.9226 {
						return 0.95181
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 6.5000 {
					if s[1] <= 0.5000 {
						return 0.93939
					} else {
						return 0.96129
					}
				} else {
					if s[5] <= 2.5944 {
						return 0.98616
					} else {
						return 0.99727
					}
				}
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[3] <= -0.7908 {
		if s[5] <= 1.7571 {
			if s[3] <= -3.4360 {
				if s[5] <= -2.7411 {
					if s[4] <= -0.6232 {
						return 0.00093
					} else {
						return 0.70000
					}
				} else {
					if s[3] <= -4.1116 {
						return 0.00000
					} else {
						return 0.08475
					}
				}
			} else {
				if s[1] <= 5.5000 {
					if s[3] <= -3.4216 {
						return 0.50000
					} else {
						return 0.00718
					}
				} else {
					if s[5] <= -4.5829 {
						return 0.31343
					} else {
						return 0.02174
					}
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[1] <= 2.5000 {
					return 0.00000
				} else {
					return 1.00000
				}
			} else {
				if s[5] <= 3.0512 {
					if s[0] <= 0.0278 {
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
		if s[6] <= -0.5292 {
			if s[4] <= 1.7858 {
				if s[0] <= 0.0823 {
					if s[1] <= 5.5000 {
						return 0.11688
					} else {
						return 0.63115
					}
				} else {
					if s[4] <= -1.1746 {
						return 0.00962
					} else {
						return 0.14062
					}
				}
			} else {
				if s[1] <= 7.5000 {
					if s[0] <= 0.0277 {
						return 0.48571
					} else {
						return 0.85714
					}
				} else {
					if s[3] <= 8.1903 {
						return 0.96809
					} else {
						return 0.60000
					}
				}
			}
		} else {
			if s[4] <= -0.4730 {
				if s[0] <= 0.2798 {
					if s[4] <= -1.2298 {
						return 0.18519
					} else {
						return 0.72414
					}
				} else {
					if s[1] <= 9.5000 {
						return 0.02469
					} else {
						return 0.33333
					}
				}
			} else {
				if s[1] <= 9.5000 {
					if s[0] <= 0.7572 {
						return 0.95542
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= 8.5054 {
						return 0.99689
					} else {
						return 0.97887
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[4] <= -0.3101 {
		if s[3] <= -2.6291 {
			if s[4] <= -1.4141 {
				if s[3] <= -6.0493 {
					if s[1] <= 9.5000 {
						return 0.00005
					} else {
						return 0.00225
					}
				} else {
					if s[3] <= -6.0438 {
						return 0.50000
					} else {
						return 0.00417
					}
				}
			} else {
				if s[6] <= -10.5375 {
					return 1.00000
				} else {
					if s[4] <= -1.3607 {
						return 1.00000
					} else {
						return 0.05455
					}
				}
			}
		} else {
			if s[1] <= 5.5000 {
				if s[5] <= 1.0451 {
					if s[2] <= 0.5000 {
						return 0.02427
					} else {
						return 0.08511
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.11828
					} else {
						return 0.60000
					}
				}
			} else {
				if s[0] <= 0.0450 {
					if s[1] <= 11.5000 {
						return 0.47826
					} else {
						return 0.16923
					}
				} else {
					if s[4] <= -0.7885 {
						return 0.06667
					} else {
						return 0.66667
					}
				}
			}
		}
	} else {
		if s[1] <= 10.5000 {
			if s[0] <= 0.2110 {
				if s[5] <= -1.0384 {
					if s[1] <= 3.5000 {
						return 0.23188
					} else {
						return 0.75238
					}
				} else {
					if s[4] <= 1.8666 {
						return 0.85141
					} else {
						return 0.96317
					}
				}
			} else {
				if s[1] <= 3.5000 {
					if s[0] <= 0.4030 {
						return 0.55963
					} else {
						return 0.25000
					}
				} else {
					if s[4] <= 1.1264 {
						return 0.42857
					} else {
						return 0.93548
					}
				}
			}
		} else {
			if s[3] <= -0.3920 {
				if s[1] <= 59.0000 {
					if s[4] <= 0.5193 {
						return 0.44444
					} else {
						return 0.90476
					}
				} else {
					return 0.00000
				}
			} else {
				if s[4] <= 2.6508 {
					if s[1] <= 83.5000 {
						return 0.91204
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= 3.8210 {
						return 0.99000
					} else {
						return 0.99873
					}
				}
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[4] <= -0.3249 {
		if s[4] <= -1.4307 {
			if s[4] <= -3.3079 {
				if s[2] <= 0.5000 {
					if s[4] <= -4.8884 {
						return 0.00000
					} else {
						return 0.00586
					}
				} else {
					if s[5] <= -11.4433 {
						return 0.00000
					} else {
						return 0.00071
					}
				}
			} else {
				if s[4] <= -3.3028 {
					return 1.00000
				} else {
					if s[3] <= -6.2435 {
						return 0.28571
					} else {
						return 0.04076
					}
				}
			}
		} else {
			if s[5] <= 0.2652 {
				if s[4] <= -1.4283 {
					return 1.00000
				} else {
					if s[1] <= 4.5000 {
						return 0.04630
					} else {
						return 0.28571
					}
				}
			} else {
				if s[3] <= 1.0544 {
					if s[2] <= 6.5000 {
						return 0.92857
					} else {
						return 0.00000
					}
				} else {
					if s[1] <= 3.5000 {
						return 0.34483
					} else {
						return 0.10526
					}
				}
			}
		}
	} else {
		if s[6] <= 0.1144 {
			if s[1] <= 4.5000 {
				if s[3] <= 2.7212 {
					if s[1] <= 3.5000 {
						return 0.02778
					} else {
						return 0.75000
					}
				} else {
					if s[3] <= 5.2574 {
						return 0.66102
					} else {
						return 0.27778
					}
				}
			} else {
				if s[3] <= -0.2813 {
					if s[4] <= 0.3351 {
						return 0.42857
					} else {
						return 0.89474
					}
				} else {
					if s[3] <= 6.5389 {
						return 0.88500
					} else {
						return 0.72414
					}
				}
			}
		} else {
			if s[4] <= 0.8432 {
				if s[3] <= 3.1122 {
					if s[0] <= 0.3106 {
						return 0.89474
					} else {
						return 0.40000
					}
				} else {
					if s[0] <= 0.3859 {
						return 0.77778
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 6.5000 {
					if s[0] <= 0.0035 {
						return 0.92971
					} else {
						return 0.97480
					}
				} else {
					if s[3] <= 5.3356 {
						return 0.99906
					} else {
						return 0.99244
					}
				}
			}
		}
	}
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[6] <= -0.5043 {
		if s[5] <= 0.9568 {
			if s[1] <= 27.5000 {
				if s[3] <= -1.2479 {
					if s[3] <= -3.1517 {
						return 0.00171
					} else {
						return 0.08046
					}
				} else {
					if s[1] <= 5.5000 {
						return 0.06122
					} else {
						return 0.58480
					}
				}
			} else {
				if s[4] <= 0.3449 {
					return 0.00000
				} else {
					if s[5] <= -5.6022 {
						return 0.95455
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[0] <= 0.1942 {
				if s[5] <= 4.6773 {
					if s[2] <= 0.5000 {
						return 0.35417
					} else {
						return 0.86066
					}
				} else {
					if s[6] <= -0.9103 {
						return 0.96450
					} else {
						return 0.50000
					}
				}
			} else {
				if s[3] <= 5.5284 {
					if s[1] <= 4.5000 {
						return 0.03571
					} else {
						return 0.17647
					}
				} else {
					if s[4] <= -1.6857 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 2.3895 {
			if s[3] <= -0.8961 {
				if s[1] <= 30.5000 {
					if s[3] <= -3.1597 {
						return 0.25000
					} else {
						return 0.02941
					}
				} else {
					if s[5] <= -5.5098 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= -0.1577 {
					if s[1] <= 9.5000 {
						return 0.36364
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.5425 {
						return 0.90630
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[0] <= 0.6938 {
				if s[1] <= 1.5000 {
					if s[6] <= 2.9708 {
						return 0.72973
					} else {
						return 0.93789
					}
				} else {
					if s[1] <= 6.5000 {
						return 0.97310
					} else {
						return 0.99564
					}
				}
			} else {
				return 0.00000
			}
		}
	}
}
