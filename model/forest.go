package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[4] <= -0.0076 {
		if s[7] <= 0.4226 {
			if s[4] <= -1.3477 {
				if s[4] <= -2.5442 {
					if s[6] <= 1.6205 {
						return 0.00050
					} else {
						return 0.05882
					}
				} else {
					if s[4] <= -2.5437 {
						return 1.00000
					} else {
						return 0.05463
					}
				}
			} else {
				if s[0] <= 0.2990 {
					if s[3] <= -0.9714 {
						return 0.06061
					} else {
						return 0.37342
					}
				} else {
					if s[1] <= 2.5000 {
						return 0.19231
					} else {
						return 0.93750
					}
				}
			}
		} else {
			if s[4] <= -5.5524 {
				if s[4] <= -6.5516 {
					return 0.00000
				} else {
					if s[7] <= 0.7750 {
						return 0.33333
					} else {
						return 0.00000
					}
				}
			} else {
				if s[8] <= 0.5000 {
					if s[3] <= -0.3212 {
						return 0.91463
					} else {
						return 0.00000
					}
				} else {
					return 0.00000
				}
			}
		}
	} else {
		if s[3] <= -0.4470 {
			if s[1] <= 4.5000 {
				if s[4] <= 1.3319 {
					if s[6] <= 0.4619 {
						return 0.01754
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= -2.7015 {
						return 0.71429
					} else {
						return 0.00000
					}
				}
			} else {
				if s[0] <= 0.0116 {
					if s[1] <= 56.5000 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[4] <= 2.5057 {
				if s[6] <= -0.4571 {
					if s[3] <= 4.3742 {
						return 0.59494
					} else {
						return 0.20000
					}
				} else {
					if s[6] <= 0.4477 {
						return 0.78082
					} else {
						return 0.94231
					}
				}
			} else {
				if s[1] <= 0.5000 {
					if s[5] <= 6.5361 {
						return 0.85965
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= 3.7994 {
						return 0.96231
					} else {
						return 0.99739
					}
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[6] <= -0.4951 {
		if s[3] <= -0.4470 {
			if s[4] <= -1.9245 {
				if s[1] <= 8.5000 {
					if s[7] <= 0.5079 {
						return 0.00009
					} else {
						return 0.24138
					}
				} else {
					if s[7] <= 0.5833 {
						return 0.00286
					} else {
						return 0.57143
					}
				}
			} else {
				if s[8] <= 0.5000 {
					if s[7] <= 0.1181 {
						return 0.04082
					} else {
						return 0.79592
					}
				} else {
					if s[6] <= -0.8030 {
						return 0.06028
					} else {
						return 0.80000
					}
				}
			}
		} else {
			if s[4] <= 0.6893 {
				if s[7] <= 0.7875 {
					if s[2] <= 0.5000 {
						return 0.04598
					} else {
						return 0.38462
					}
				} else {
					return 1.00000
				}
			} else {
				if s[2] <= 0.5000 {
					if s[8] <= 0.5000 {
						return 1.00000
					} else {
						return 0.40000
					}
				} else {
					if s[4] <= 2.3153 {
						return 0.89474
					} else {
						return 0.98684
					}
				}
			}
		}
	} else {
		if s[0] <= 0.5603 {
			if s[3] <= -0.4470 {
				if s[4] <= -0.9878 {
					return 0.00000
				} else {
					if s[3] <= -2.7096 {
						return 1.00000
					} else {
						return 0.30303
					}
				}
			} else {
				if s[1] <= 9.5000 {
					if s[6] <= 2.1948 {
						return 0.83881
					} else {
						return 0.96982
					}
				} else {
					if s[3] <= -0.3853 {
						return 0.75000
					} else {
						return 0.99519
					}
				}
			}
		} else {
			if s[4] <= -3.1567 {
				return 0.00000
			} else {
				if s[4] <= -2.6836 {
					return 1.00000
				} else {
					return 0.00000
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[4] <= -0.4302 {
		if s[6] <= -0.5688 {
			if s[0] <= 0.1132 {
				if s[4] <= -2.1575 {
					if s[7] <= 0.4143 {
						return 0.00163
					} else {
						return 0.85714
					}
				} else {
					if s[5] <= -6.1791 {
						return 0.28571
					} else {
						return 0.04930
					}
				}
			} else {
				if s[6] <= -1.2999 {
					if s[0] <= 0.1509 {
						return 0.00581
					} else {
						return 0.00005
					}
				} else {
					if s[5] <= -2.0585 {
						return 0.00000
					} else {
						return 0.28571
					}
				}
			}
		} else {
			if s[5] <= -1.4495 {
				if s[5] <= -6.3180 {
					if s[4] <= -1.6965 {
						return 0.00000
					} else {
						return 0.50000
					}
				} else {
					return 0.00000
				}
			} else {
				if s[4] <= -1.2728 {
					if s[3] <= -0.1500 {
						return 0.00000
					} else {
						return 0.14286
					}
				} else {
					if s[4] <= -0.6227 {
						return 0.91176
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.4596 {
			if s[2] <= 0.5000 {
				if s[8] <= 0.5000 {
					if s[7] <= 0.0909 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.2335 {
						return 0.35385
					} else {
						return 0.21667
					}
				}
			} else {
				if s[3] <= -0.5317 {
					if s[6] <= -6.8678 {
						return 0.57143
					} else {
						return 0.14286
					}
				} else {
					if s[4] <= 0.6305 {
						return 0.57778
					} else {
						return 0.96369
					}
				}
			}
		} else {
			if s[4] <= 0.1297 {
				if s[1] <= 22.0000 {
					if s[3] <= 1.3678 {
						return 1.00000
					} else {
						return 0.51852
					}
				} else {
					return 0.00000
				}
			} else {
				if s[3] <= -0.5317 {
					if s[0] <= 0.1406 {
						return 0.81250
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= -10.2267 {
						return 0.00000
					} else {
						return 0.98752
					}
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[4] <= -0.3933 {
		if s[5] <= -1.1107 {
			if s[4] <= -1.9245 {
				if s[7] <= 0.5079 {
					if s[6] <= 5.6039 {
						return 0.00050
					} else {
						return 0.50000
					}
				} else {
					if s[0] <= 0.5000 {
						return 0.96000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[5] <= -6.3373 {
					if s[7] <= 0.7071 {
						return 0.19310
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= -1.5247 {
						return 0.00694
					} else {
						return 0.20000
					}
				}
			}
		} else {
			if s[4] <= -1.8201 {
				if s[5] <= -1.0696 {
					return 1.00000
				} else {
					if s[1] <= 17.0000 {
						return 0.02538
					} else {
						return 0.20000
					}
				}
			} else {
				if s[1] <= 6.5000 {
					if s[6] <= -0.9599 {
						return 0.00000
					} else {
						return 0.57576
					}
				} else {
					if s[3] <= -0.3346 {
						return 0.50000
					} else {
						return 0.83333
					}
				}
			}
		}
	} else {
		if s[1] <= 4.5000 {
			if s[3] <= -0.0057 {
				if s[3] <= -0.4470 {
					if s[5] <= 0.4710 {
						return 0.03125
					} else {
						return 0.25000
					}
				} else {
					if s[0] <= 0.0543 {
						return 0.00000
					} else {
						return 0.80000
					}
				}
			} else {
				if s[0] <= 0.0055 {
					if s[6] <= 0.5832 {
						return 0.28788
					} else {
						return 0.90830
					}
				} else {
					if s[6] <= -0.1404 {
						return 0.34615
					} else {
						return 0.95762
					}
				}
			}
		} else {
			if s[3] <= 0.6721 {
				if s[4] <= 1.3014 {
					if s[4] <= 1.1218 {
						return 0.68750
					} else {
						return 0.14286
					}
				} else {
					if s[0] <= 0.0549 {
						return 1.00000
					} else {
						return 0.94737
					}
				}
			} else {
				if s[6] <= 0.0486 {
					if s[6] <= 0.0074 {
						return 0.91977
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= 2.5646 {
						return 0.93750
					} else {
						return 1.00000
					}
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[6] <= -0.5551 {
		if s[4] <= 1.2123 {
			if s[4] <= -2.1331 {
				if s[4] <= -4.8611 {
					if s[7] <= 0.5357 {
						return 0.00013
					} else {
						return 0.05882
					}
				} else {
					if s[7] <= 0.4143 {
						return 0.00887
					} else {
						return 0.93548
					}
				}
			} else {
				if s[7] <= 0.2361 {
					if s[2] <= 0.5000 {
						return 0.05974
					} else {
						return 0.28571
					}
				} else {
					if s[4] <= -0.0944 {
						return 0.92857
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[1] <= 1.5000 {
				if s[0] <= 0.0091 {
					if s[5] <= 2.3114 {
						return 0.12500
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.0597 {
						return 1.00000
					} else {
						return 0.16667
					}
				}
			} else {
				if s[0] <= 0.2225 {
					if s[3] <= 9.5092 {
						return 0.95362
					} else {
						return 0.50000
					}
				} else {
					if s[1] <= 8.5000 {
						return 0.20000
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[4] <= -0.4145 {
			if s[3] <= -0.7284 {
				if s[4] <= -0.8059 {
					return 0.00000
				} else {
					if s[5] <= -3.2086 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= 4.2220 {
					if s[5] <= -5.6826 {
						return 0.73333
					} else {
						return 0.36486
					}
				} else {
					if s[3] <= 2.4503 {
						return 0.22222
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[5] <= 0.3274 {
				if s[4] <= 0.0302 {
					if s[3] <= 2.1682 {
						return 0.77778
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= 0.2942 {
						return 0.93134
					} else {
						return 0.00000
					}
				}
			} else {
				if s[4] <= 2.5321 {
					if s[7] <= 0.1500 {
						return 0.92939
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= -0.5434 {
						return 0.00000
					} else {
						return 0.99684
					}
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[6] <= -0.4113 {
		if s[3] <= -0.6976 {
			if s[3] <= -2.1575 {
				if s[7] <= 0.4226 {
					if s[2] <= 0.5000 {
						return 0.00014
					} else {
						return 0.00469
					}
				} else {
					if s[4] <= -5.2000 {
						return 0.06818
					} else {
						return 0.94000
					}
				}
			} else {
				if s[7] <= 0.7803 {
					if s[1] <= 5.5000 {
						return 0.00000
					} else {
						return 0.19588
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[8] <= 0.5000 {
					if s[1] <= 3.5000 {
						return 0.00000
					} else {
						return 0.85185
					}
				} else {
					if s[0] <= 0.2178 {
						return 0.31481
					} else {
						return 0.03763
					}
				}
			} else {
				if s[0] <= 0.2060 {
					if s[1] <= 8.5000 {
						return 0.63265
					} else {
						return 0.93168
					}
				} else {
					if s[4] <= -0.3496 {
						return 0.00000
					} else {
						return 0.50000
					}
				}
			}
		}
	} else {
		if s[4] <= -0.5597 {
			if s[6] <= 1.4489 {
				if s[3] <= 1.6126 {
					if s[6] <= -0.2738 {
						return 0.20000
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 3.0934 {
						return 1.00000
					} else {
						return 0.10000
					}
				}
			} else {
				if s[3] <= 3.9453 {
					if s[6] <= 2.0642 {
						return 0.86667
					} else {
						return 0.29508
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[6] <= 0.6165 {
				if s[3] <= -0.1589 {
					if s[5] <= 1.1157 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= 0.3923 {
						return 0.50000
					} else {
						return 0.89109
					}
				}
			} else {
				if s[3] <= -0.7427 {
					if s[3] <= -1.5050 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.7258 {
						return 0.98461
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[4] <= -0.3933 {
		if s[3] <= -1.0999 {
			if s[3] <= -2.9326 {
				if s[7] <= 0.5357 {
					if s[4] <= -6.2109 {
						return 0.00000
					} else {
						return 0.00462
					}
				} else {
					if s[8] <= 0.5000 {
						return 0.69697
					} else {
						return 0.00000
					}
				}
			} else {
				if s[6] <= 1.6589 {
					if s[8] <= 0.5000 {
						return 0.28070
					} else {
						return 0.01828
					}
				} else {
					if s[4] <= -1.8162 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[4] <= -2.7474 {
				if s[1] <= 12.5000 {
					if s[1] <= 0.5000 {
						return 0.01852
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 1.6617 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= -1.8043 {
					if s[8] <= 0.5000 {
						return 0.40000
					} else {
						return 0.10909
					}
				} else {
					if s[3] <= 3.1037 {
						return 0.55319
					} else {
						return 0.15625
					}
				}
			}
		}
	} else {
		if s[6] <= 0.0410 {
			if s[1] <= 5.5000 {
				if s[4] <= 1.6548 {
					if s[7] <= 0.3000 {
						return 0.10317
					} else {
						return 1.00000
					}
				} else {
					if s[3] <= 5.4852 {
						return 0.76471
					} else {
						return 0.10000
					}
				}
			} else {
				if s[5] <= 1.7741 {
					if s[0] <= 0.0139 {
						return 0.89216
					} else {
						return 0.45455
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.77273
					} else {
						return 0.94167
					}
				}
			}
		} else {
			if s[6] <= 2.5550 {
				if s[6] <= 2.5535 {
					if s[5] <= 2.2562 {
						return 0.85546
					} else {
						return 0.97471
					}
				} else {
					return 0.00000
				}
			} else {
				if s[1] <= 0.5000 {
					if s[3] <= 6.2413 {
						return 0.91667
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.7304 {
						return 0.99304
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[4] <= -0.4040 {
		if s[5] <= -1.3340 {
			if s[1] <= 4.5000 {
				if s[1] <= 2.5000 {
					if s[4] <= -2.7978 {
						return 0.00000
					} else {
						return 0.01916
					}
				} else {
					if s[3] <= 3.0158 {
						return 0.00257
					} else {
						return 0.75000
					}
				}
			} else {
				if s[8] <= 0.5000 {
					if s[0] <= 0.0516 {
						return 0.27053
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= -2.0340 {
						return 0.00024
					} else {
						return 0.10185
					}
				}
			}
		} else {
			if s[6] <= 1.2471 {
				if s[7] <= 0.0556 {
					if s[1] <= 2.5000 {
						return 0.03636
					} else {
						return 0.15094
					}
				} else {
					return 1.00000
				}
			} else {
				if s[6] <= 1.9211 {
					if s[4] <= -0.9391 {
						return 1.00000
					} else {
						return 0.50000
					}
				} else {
					if s[4] <= -2.3621 {
						return 0.04545
					} else {
						return 0.60870
					}
				}
			}
		}
	} else {
		if s[6] <= 0.1111 {
			if s[2] <= 0.5000 {
				if s[7] <= 0.3571 {
					if s[5] <= -0.5367 {
						return 0.08571
					} else {
						return 0.45385
					}
				} else {
					if s[3] <= -0.0968 {
						return 0.81818
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= 1.8475 {
					if s[3] <= 3.2445 {
						return 0.74286
					} else {
						return 0.25000
					}
				} else {
					if s[3] <= 3.9823 {
						return 0.98726
					} else {
						return 0.92568
					}
				}
			}
		} else {
			if s[4] <= 1.2971 {
				if s[5] <= -1.8841 {
					if s[4] <= 0.7714 {
						return 0.73333
					} else {
						return 0.12500
					}
				} else {
					if s[0] <= 0.1983 {
						return 0.93985
					} else {
						return 0.75581
					}
				}
			} else {
				if s[6] <= 2.1800 {
					if s[1] <= 9.5000 {
						return 0.84865
					} else {
						return 0.98897
					}
				} else {
					if s[5] <= 0.5888 {
						return 0.96016
					} else {
						return 0.99453
					}
				}
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[3] <= -0.4512 {
		if s[4] <= 1.0970 {
			if s[1] <= 6.5000 {
				if s[5] <= 0.6503 {
					if s[7] <= 0.4018 {
						return 0.00039
					} else {
						return 0.50000
					}
				} else {
					if s[4] <= 0.8463 {
						return 0.03571
					} else {
						return 1.00000
					}
				}
			} else {
				if s[5] <= -0.6645 {
					if s[0] <= 0.0132 {
						return 0.05335
					} else {
						return 0.00040
					}
				} else {
					if s[2] <= 0.5000 {
						return 1.00000
					} else {
						return 0.37500
					}
				}
			}
		} else {
			if s[1] <= 5.0000 {
				if s[5] <= -2.7175 {
					if s[4] <= 4.7159 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					return 0.00000
				}
			} else {
				if s[8] <= 0.5000 {
					return 1.00000
				} else {
					if s[0] <= 0.0278 {
						return 1.00000
					} else {
						return 0.25000
					}
				}
			}
		}
	} else {
		if s[4] <= 0.0693 {
			if s[4] <= -1.3551 {
				if s[3] <= 0.7487 {
					if s[4] <= -2.8002 {
						return 0.07407
					} else {
						return 0.50000
					}
				} else {
					if s[6] <= 0.5905 {
						return 0.00513
					} else {
						return 0.11111
					}
				}
			} else {
				if s[3] <= 2.9536 {
					if s[5] <= 3.5976 {
						return 0.69231
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.2724 {
						return 0.11364
					} else {
						return 0.61905
					}
				}
			}
		} else {
			if s[6] <= 0.4477 {
				if s[4] <= 2.7563 {
					if s[2] <= 0.5000 {
						return 0.47260
					} else {
						return 0.83582
					}
				} else {
					if s[4] <= 3.4529 {
						return 0.88889
					} else {
						return 0.97967
					}
				}
			} else {
				if s[4] <= 2.6768 {
					if s[6] <= 8.7564 {
						return 0.95434
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 1.1641 {
						return 0.93651
					} else {
						return 0.99556
					}
				}
			}
		}
	}
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[1] <= 15.5000 {
		if s[6] <= -0.3564 {
			if s[4] <= 0.6974 {
				if s[3] <= -3.1362 {
					if s[0] <= 0.0064 {
						return 0.00805
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= -1.9283 {
						return 0.02930
					} else {
						return 0.15152
					}
				}
			} else {
				if s[7] <= 0.3636 {
					if s[2] <= 0.5000 {
						return 0.36111
					} else {
						return 0.77108
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[0] <= 0.5708 {
				if s[5] <= 0.8834 {
					if s[3] <= -0.4470 {
						return 0.10714
					} else {
						return 0.85294
					}
				} else {
					if s[6] <= 2.0682 {
						return 0.82700
					} else {
						return 0.96690
					}
				}
			} else {
				if s[1] <= 12.5000 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		}
	} else {
		if s[5] <= -2.4089 {
			if s[4] <= 0.8732 {
				if s[3] <= -0.6852 {
					if s[4] <= -6.1833 {
						return 0.00000
					} else {
						return 0.01345
					}
				} else {
					if s[4] <= -0.6007 {
						return 0.00000
					} else {
						return 0.76923
					}
				}
			} else {
				if s[3] <= -0.6663 {
					if s[1] <= 17.0000 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[7] <= 0.1550 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[4] <= -0.0762 {
				if s[6] <= -2.7601 {
					if s[5] <= 2.5765 {
						return 0.00000
					} else {
						return 0.25000
					}
				} else {
					if s[6] <= -0.2972 {
						return 0.87500
					} else {
						return 0.27273
					}
				}
			} else {
				if s[4] <= 3.3662 {
					if s[4] <= 3.3643 {
						return 0.95816
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 1.1641 {
						return 0.97561
					} else {
						return 0.99899
					}
				}
			}
		}
	}
}
