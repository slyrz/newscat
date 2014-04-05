package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[3] <= -0.3898 {
		if s[7] <= 0.4808 {
			if s[6] <= 3.2427 {
				if s[5] <= -0.4061 {
					if s[0] <= 0.0735 {
						return 0.01088
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= -0.8299 {
						return 0.00000
					} else {
						return 0.35556
					}
				}
			} else {
				return 1.00000
			}
		} else {
			if s[3] <= -9.7990 {
				return 0.00000
			} else {
				if s[0] <= 0.5000 {
					if s[1] <= 9.5000 {
						return 0.97619
					} else {
						return 0.66667
					}
				} else {
					return 0.00000
				}
			}
		}
	} else {
		if s[1] <= 10.5000 {
			if s[6] <= -0.3930 {
				if s[0] <= 0.2064 {
					if s[1] <= 6.5000 {
						return 0.44144
					} else {
						return 0.84000
					}
				} else {
					if s[5] <= 0.5082 {
						return 0.02326
					} else {
						return 0.23913
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[4] <= -0.3409 {
						return 0.27174
					} else {
						return 0.94289
					}
				} else {
					if s[4] <= 2.6588 {
						return 0.92405
					} else {
						return 0.99548
					}
				}
			}
		} else {
			if s[4] <= 0.6627 {
				if s[3] <= 2.5259 {
					if s[3] <= 0.5539 {
						return 0.50000
					} else {
						return 0.90476
					}
				} else {
					if s[0] <= 0.3627 {
						return 0.00000
					} else {
						return 0.33333
					}
				}
			} else {
				if s[4] <= 2.1072 {
					if s[1] <= 76.0000 {
						return 0.95238
					} else {
						return 0.00000
					}
				} else {
					if s[2] <= 5.5000 {
						return 0.99911
					} else {
						return 0.98413
					}
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[5] <= 0.4059 {
		if s[6] <= 0.4255 {
			if s[4] <= -1.5895 {
				if s[7] <= 0.5278 {
					if s[4] <= -2.4340 {
						return 0.00040
					} else {
						return 0.04455
					}
				} else {
					if s[0] <= 0.5000 {
						return 0.71429
					} else {
						return 0.00000
					}
				}
			} else {
				if s[7] <= 0.1847 {
					if s[6] <= -0.8362 {
						return 0.19142
					} else {
						return 0.53333
					}
				} else {
					if s[7] <= 0.8258 {
						return 0.66667
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[0] <= 0.2091 {
				if s[1] <= 10.5000 {
					if s[3] <= -0.7206 {
						return 0.13333
					} else {
						return 0.80357
					}
				} else {
					if s[3] <= -0.7853 {
						return 0.50000
					} else {
						return 0.99359
					}
				}
			} else {
				if s[4] <= 1.2962 {
					if s[5] <= -1.8816 {
						return 0.03333
					} else {
						return 0.23529
					}
				} else {
					if s[6] <= 5.8251 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[4] <= -0.4101 {
			if s[0] <= 0.1963 {
				if s[1] <= 10.0000 {
					if s[6] <= -0.8081 {
						return 0.09091
					} else {
						return 0.66667
					}
				} else {
					if s[2] <= 6.5000 {
						return 0.90909
					} else {
						return 0.00000
					}
				}
			} else {
				if s[0] <= 0.4468 {
					if s[3] <= 1.2469 {
						return 0.00000
					} else {
						return 0.22222
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[6] <= 1.8802 {
				if s[4] <= 2.9123 {
					if s[1] <= 4.5000 {
						return 0.67544
					} else {
						return 0.89143
					}
				} else {
					if s[3] <= 5.5768 {
						return 0.99301
					} else {
						return 0.97368
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[0] <= 0.7203 {
						return 0.96815
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= 2.6550 {
						return 0.93750
					} else {
						return 0.99946
					}
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[3] <= -0.3898 {
		if s[7] <= 0.4808 {
			if s[4] <= 1.3123 {
				if s[3] <= -2.2372 {
					if s[3] <= -3.1543 {
						return 0.00035
					} else {
						return 0.01832
					}
				} else {
					if s[1] <= 11.5000 {
						return 0.06103
					} else {
						return 0.27778
					}
				}
			} else {
				if s[4] <= 2.8733 {
					if s[3] <= -0.6844 {
						return 0.35714
					} else {
						return 1.00000
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[4] <= -4.2949 {
				if s[0] <= 0.5000 {
					return 1.00000
				} else {
					return 0.00000
				}
			} else {
				if s[3] <= -1.0714 {
					if s[7] <= 0.8258 {
						return 0.89655
					} else {
						return 1.00000
					}
				} else {
					if s[7] <= 0.7667 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.4073 {
			if s[4] <= 0.5959 {
				if s[3] <= 2.3761 {
					if s[4] <= -1.5681 {
						return 0.17949
					} else {
						return 0.78261
					}
				} else {
					if s[0] <= 0.2649 {
						return 0.24528
					} else {
						return 0.68750
					}
				}
			} else {
				if s[6] <= -0.4154 {
					if s[4] <= 2.7793 {
						return 0.57143
					} else {
						return 0.95385
					}
				} else {
					if s[4] <= 2.6443 {
						return 0.92349
					} else {
						return 0.99495
					}
				}
			}
		} else {
			if s[1] <= 11.5000 {
				if s[0] <= 0.5708 {
					if s[3] <= 2.3167 {
						return 0.00000
					} else {
						return 0.56000
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.00000
					} else {
						return 0.06667
					}
				}
			} else {
				if s[5] <= -4.2351 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[5] <= -0.5964 {
		if s[6] <= -0.6269 {
			if s[3] <= -2.0340 {
				if s[4] <= 0.9813 {
					if s[4] <= -4.5129 {
						return 0.00021
					} else {
						return 0.03208
					}
				} else {
					if s[6] <= -8.7579 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[0] <= 0.0099 {
					if s[3] <= 1.7213 {
						return 0.27604
					} else {
						return 0.73214
					}
				} else {
					if s[6] <= -1.3400 {
						return 0.02454
					} else {
						return 0.22222
					}
				}
			}
		} else {
			if s[5] <= -4.2192 {
				if s[3] <= -2.2554 {
					return 0.00000
				} else {
					if s[0] <= 0.3577 {
						return 0.93827
					} else {
						return 0.00000
					}
				}
			} else {
				if s[5] <= -1.1548 {
					if s[2] <= 0.5000 {
						return 0.02857
					} else {
						return 0.80000
					}
				} else {
					if s[1] <= 7.0000 {
						return 0.61538
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.4073 {
			if s[4] <= -0.4089 {
				if s[1] <= 2.5000 {
					if s[0] <= 0.1289 {
						return 0.08824
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= -1.8728 {
						return 0.00000
					} else {
						return 0.46341
					}
				}
			} else {
				if s[4] <= 1.9289 {
					if s[5] <= 6.8058 {
						return 0.81910
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= 0.8160 {
						return 0.93248
					} else {
						return 0.99537
					}
				}
			}
		} else {
			if s[6] <= -0.9425 {
				return 0.00000
			} else {
				if s[0] <= 0.7090 {
					if s[4] <= -0.8815 {
						return 0.12500
					} else {
						return 0.61538
					}
				} else {
					return 0.00000
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[3] <= -0.3898 {
		if s[6] <= 3.2604 {
			if s[7] <= 0.6833 {
				if s[3] <= -2.0332 {
					if s[6] <= 0.7466 {
						return 0.00054
					} else {
						return 0.40000
					}
				} else {
					if s[7] <= 0.1225 {
						return 0.11297
					} else {
						return 0.81818
					}
				}
			} else {
				if s[0] <= 0.5000 {
					if s[3] <= -3.4434 {
						return 1.00000
					} else {
						return 0.71429
					}
				} else {
					return 0.00000
				}
			}
		} else {
			return 1.00000
		}
	} else {
		if s[5] <= 0.5530 {
			if s[1] <= 7.5000 {
				if s[4] <= 0.5809 {
					if s[0] <= 0.1856 {
						return 0.27586
					} else {
						return 0.02759
					}
				} else {
					if s[4] <= 1.9914 {
						return 0.52273
					} else {
						return 0.86538
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[3] <= 2.2321 {
						return 0.17647
					} else {
						return 0.80000
					}
				} else {
					if s[1] <= 75.0000 {
						return 0.98707
					} else {
						return 0.75000
					}
				}
			}
		} else {
			if s[6] <= -0.5506 {
				if s[0] <= 0.2064 {
					if s[4] <= 0.6423 {
						return 0.35714
					} else {
						return 0.93145
					}
				} else {
					if s[6] <= -1.0865 {
						return 0.05172
					} else {
						return 0.75000
					}
				}
			} else {
				if s[1] <= 10.5000 {
					if s[0] <= 0.3376 {
						return 0.96519
					} else {
						return 0.24324
					}
				} else {
					if s[4] <= 0.0784 {
						return 0.00000
					} else {
						return 0.99687
					}
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[3] <= -0.3898 {
		if s[4] <= -1.9804 {
			if s[7] <= 0.7071 {
				if s[0] <= 0.0121 {
					if s[1] <= 8.5000 {
						return 0.00059
					} else {
						return 0.01134
					}
				} else {
					return 0.00000
				}
			} else {
				if s[8] <= 0.5000 {
					if s[4] <= -3.0107 {
						return 0.61538
					} else {
						return 1.00000
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[1] <= 5.5000 {
				if s[6] <= -5.7044 {
					if s[4] <= 1.2684 {
						return 0.00775
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= -5.5731 {
						return 1.00000
					} else {
						return 0.01299
					}
				}
			} else {
				if s[8] <= 0.5000 {
					if s[7] <= 0.3056 {
						return 0.28571
					} else {
						return 0.84615
					}
				} else {
					if s[4] <= 1.8637 {
						return 0.17391
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.5708 {
			if s[6] <= -0.3493 {
				if s[2] <= 0.5000 {
					if s[1] <= 5.5000 {
						return 0.15517
					} else {
						return 0.64407
					}
				} else {
					if s[0] <= 0.2060 {
						return 0.94529
					} else {
						return 0.20000
					}
				}
			} else {
				if s[5] <= 1.5480 {
					if s[5] <= -3.5878 {
						return 0.97500
					} else {
						return 0.86731
					}
				} else {
					if s[0] <= 0.2317 {
						return 0.98954
					} else {
						return 0.85577
					}
				}
			}
		} else {
			if s[6] <= 5.5742 {
				return 0.00000
			} else {
				if s[6] <= 5.6275 {
					return 1.00000
				} else {
					return 0.00000
				}
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[1] <= 15.5000 {
		if s[4] <= -0.4175 {
			if s[7] <= 0.5000 {
				if s[7] <= 0.2330 {
					if s[5] <= -1.0016 {
						return 0.00130
					} else {
						return 0.11739
					}
				} else {
					if s[6] <= -7.5309 {
						return 0.00000
					} else {
						return 0.75000
					}
				}
			} else {
				if s[0] <= 0.5000 {
					if s[3] <= -3.1703 {
						return 0.97222
					} else {
						return 0.57143
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[6] <= -0.4043 {
				if s[1] <= 4.5000 {
					if s[4] <= 1.9739 {
						return 0.14545
					} else {
						return 0.48387
					}
				} else {
					if s[1] <= 10.5000 {
						return 0.67778
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= 0.8190 {
					if s[5] <= -0.2455 {
						return 0.21429
					} else {
						return 0.79808
					}
				} else {
					if s[4] <= 2.7035 {
						return 0.92041
					} else {
						return 0.99067
					}
				}
			}
		}
	} else {
		if s[3] <= -0.8298 {
			if s[6] <= -0.3383 {
				if s[4] <= -2.4415 {
					if s[8] <= 0.5000 {
						return 0.04000
					} else {
						return 0.00000
					}
				} else {
					if s[7] <= 0.0938 {
						return 0.11364
					} else {
						return 1.00000
					}
				}
			} else {
				if s[1] <= 34.5000 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		} else {
			if s[6] <= -1.5681 {
				if s[4] <= 0.8731 {
					if s[3] <= 2.1863 {
						return 0.71429
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= -1.9504 {
						return 0.99505
					} else {
						return 0.83333
					}
				}
			} else {
				if s[3] <= -0.6844 {
					if s[4] <= 0.9798 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= 2.8682 {
						return 0.98466
					} else {
						return 1.00000
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[3] <= -0.3898 {
		if s[7] <= 0.4226 {
			if s[3] <= -2.0213 {
				if s[4] <= 0.8722 {
					if s[1] <= 45.5000 {
						return 0.00040
					} else {
						return 0.16667
					}
				} else {
					if s[0] <= 0.0291 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= 3.2612 {
					if s[5] <= 1.3274 {
						return 0.12987
					} else {
						return 0.50000
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[8] <= 0.5000 {
				if s[3] <= -11.0252 {
					return 0.00000
				} else {
					if s[2] <= 0.5000 {
						return 0.69388
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
			if s[4] <= -0.1832 {
				if s[4] <= -1.1796 {
					if s[4] <= -2.3168 {
						return 0.00658
					} else {
						return 0.13043
					}
				} else {
					if s[3] <= 3.1057 {
						return 0.51852
					} else {
						return 0.04545
					}
				}
			} else {
				if s[0] <= 0.3308 {
					if s[6] <= -0.3769 {
						return 0.57426
					} else {
						return 0.95862
					}
				} else {
					if s[5] <= 3.8055 {
						return 0.37500
					} else {
						return 0.90000
					}
				}
			}
		} else {
			if s[4] <= -0.6066 {
				if s[0] <= 0.2059 {
					if s[6] <= 2.0075 {
						return 0.78571
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= 1.1839 {
						return 0.00000
					} else {
						return 0.20000
					}
				}
			} else {
				if s[6] <= -1.5961 {
					if s[5] <= 4.1357 {
						return 0.85034
					} else {
						return 0.96026
					}
				} else {
					if s[4] <= 3.1998 {
						return 0.98220
					} else {
						return 1.00000
					}
				}
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[5] <= -0.5979 {
		if s[6] <= 0.4703 {
			if s[6] <= -2.2273 {
				if s[3] <= -1.9015 {
					if s[3] <= -4.6318 {
						return 0.00067
					} else {
						return 0.02691
					}
				} else {
					if s[4] <= 1.9851 {
						return 0.13953
					} else {
						return 1.00000
					}
				}
			} else {
				if s[1] <= 2.5000 {
					if s[5] <= -0.8575 {
						return 0.00000
					} else {
						return 0.33333
					}
				} else {
					if s[1] <= 4.5000 {
						return 0.61538
					} else {
						return 0.19444
					}
				}
			}
		} else {
			if s[4] <= -0.7272 {
				if s[1] <= 11.5000 {
					if s[6] <= 0.8572 {
						return 0.40000
					} else {
						return 0.00000
					}
				} else {
					if s[1] <= 14.0000 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 6.5000 {
					if s[5] <= -1.9758 {
						return 0.93103
					} else {
						return 0.16667
					}
				} else {
					if s[4] <= 2.5455 {
						return 0.88889
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[5] <= 1.2281 {
			if s[1] <= 9.5000 {
				if s[5] <= 0.2605 {
					if s[1] <= 0.5000 {
						return 0.00000
					} else {
						return 0.47778
					}
				} else {
					if s[2] <= 1.5000 {
						return 0.65079
					} else {
						return 0.00000
					}
				}
			} else {
				if s[3] <= -0.6844 {
					if s[5] <= 0.8414 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[2] <= 2.5000 {
						return 0.98718
					} else {
						return 0.66667
					}
				}
			}
		} else {
			if s[4] <= -0.3362 {
				if s[1] <= 10.5000 {
					if s[0] <= 0.3101 {
						return 0.36842
					} else {
						return 0.00000
					}
				} else {
					if s[2] <= 2.5000 {
						return 0.66667
					} else {
						return 0.00000
					}
				}
			} else {
				if s[4] <= 1.9719 {
					if s[5] <= 7.6079 {
						return 0.82051
					} else {
						return 0.00000
					}
				} else {
					if s[7] <= 0.1714 {
						return 0.99104
					} else {
						return 0.50000
					}
				}
			}
		}
	}
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[6] <= -0.5414 {
		if s[4] <= 1.1964 {
			if s[7] <= 0.4808 {
				if s[4] <= -1.7550 {
					if s[4] <= -1.9095 {
						return 0.00069
					} else {
						return 0.10256
					}
				} else {
					if s[1] <= 6.5000 {
						return 0.06897
					} else {
						return 0.47101
					}
				}
			} else {
				if s[4] <= -4.6452 {
					if s[6] <= -10.3267 {
						return 1.00000
					} else {
						return 0.05882
					}
				} else {
					if s[4] <= -0.4575 {
						return 0.71429
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[1] <= 4.5000 {
				if s[4] <= 3.0599 {
					if s[3] <= 4.1084 {
						return 0.45946
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= 5.1006 {
						return 1.00000
					} else {
						return 0.50000
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[0] <= 0.2258 {
						return 0.87179
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= 3.2433 {
						return 0.96970
					} else {
						return 0.99401
					}
				}
			}
		}
	} else {
		if s[4] <= -0.0149 {
			if s[5] <= -0.5817 {
				if s[8] <= 0.5000 {
					if s[0] <= 0.0516 {
						return 0.10000
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.2106 {
						return 0.20000
					} else {
						return 0.03571
					}
				}
			} else {
				if s[4] <= -1.9330 {
					if s[5] <= 4.6303 {
						return 0.00000
					} else {
						return 0.10000
					}
				} else {
					if s[3] <= 3.9374 {
						return 0.83333
					} else {
						return 0.22222
					}
				}
			}
		} else {
			if s[4] <= 2.3308 {
				if s[3] <= -0.6991 {
					if s[3] <= -1.2240 {
						return 0.66667
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 3.1796 {
						return 0.94909
					} else {
						return 0.81961
					}
				}
			} else {
				if s[3] <= -0.9423 {
					if s[5] <= 0.7459 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.2313 {
						return 0.99539
					} else {
						return 0.94643
					}
				}
			}
		}
	}
}
