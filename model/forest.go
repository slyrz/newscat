package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[3] <= -0.3862 {
		if s[7] <= 0.5278 {
			if s[6] <= 1.2087 {
				if s[4] <= -1.2038 {
					if s[3] <= -2.3663 {
						return 0.00024
					} else {
						return 0.04396
					}
				} else {
					if s[3] <= -1.1462 {
						return 0.03333
					} else {
						return 0.23438
					}
				}
			} else {
				if s[6] <= 1.2723 {
					return 1.00000
				} else {
					if s[6] <= 2.7304 {
						return 0.10000
					} else {
						return 0.87500
					}
				}
			}
		} else {
			if s[4] <= -4.2522 {
				return 0.00000
			} else {
				if s[7] <= 0.8819 {
					if s[8] <= 0.5000 {
						return 0.91667
					} else {
						return 0.00000
					}
				} else {
					return 1.00000
				}
			}
		}
	} else {
		if s[0] <= 0.5708 {
			if s[0] <= 0.2111 {
				if s[6] <= 0.4546 {
					if s[3] <= 9.5481 {
						return 0.79129
					} else {
						return 0.00000
					}
				} else {
					if s[1] <= 5.5000 {
						return 0.94768
					} else {
						return 0.99327
					}
				}
			} else {
				if s[4] <= 0.2819 {
					if s[1] <= 3.5000 {
						return 0.05455
					} else {
						return 0.63636
					}
				} else {
					if s[6] <= 0.0169 {
						return 0.35294
					} else {
						return 0.94231
					}
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[4] <= -1.7948 {
					return 0.00000
				} else {
					if s[5] <= -7.2926 {
						return 1.00000
					} else {
						return 0.09091
					}
				}
			} else {
				if s[1] <= 2.0000 {
					if s[4] <= -2.6049 {
						return 0.00000
					} else {
						return 0.83333
					}
				} else {
					return 0.00000
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[3] <= -0.3947 {
		if s[3] <= -2.1971 {
			if s[7] <= 0.5085 {
				if s[3] <= -2.8281 {
					if s[8] <= 0.5000 {
						return 0.00000
					} else {
						return 0.00008
					}
				} else {
					if s[5] <= 0.3348 {
						return 0.00410
					} else {
						return 0.41667
					}
				}
			} else {
				if s[3] <= -7.8933 {
					if s[3] <= -8.9144 {
						return 0.00000
					} else {
						return 0.33333
					}
				} else {
					if s[0] <= 0.5000 {
						return 0.90698
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[7] <= 0.4654 {
				if s[1] <= 8.5000 {
					if s[6] <= 2.0294 {
						return 0.01031
					} else {
						return 0.40000
					}
				} else {
					if s[6] <= 1.1804 {
						return 0.32787
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= -1.6766 {
					if s[4] <= -1.7941 {
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
		if s[2] <= 0.5000 {
			if s[1] <= 0.5000 {
				if s[0] <= 0.1925 {
					if s[3] <= 2.0005 {
						return 0.35714
					} else {
						return 0.85714
					}
				} else {
					if s[4] <= 0.3710 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= -0.5957 {
					if s[7] <= 0.1667 {
						return 0.20243
					} else {
						return 0.85714
					}
				} else {
					if s[0] <= 0.5315 {
						return 0.94863
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[6] <= 0.0795 {
				if s[5] <= 0.5716 {
					if s[3] <= 3.1659 {
						return 0.53333
					} else {
						return 0.84211
					}
				} else {
					if s[0] <= 0.1921 {
						return 0.91126
					} else {
						return 0.21053
					}
				}
			} else {
				if s[4] <= -0.0099 {
					if s[5] <= -0.5010 {
						return 0.00000
					} else {
						return 0.80000
					}
				} else {
					if s[5] <= 3.6844 {
						return 0.99020
					} else {
						return 0.99800
					}
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[4] <= -0.2744 {
		if s[5] <= -1.1648 {
			if s[3] <= -2.1395 {
				if s[4] <= -1.1941 {
					if s[7] <= 0.4143 {
						return 0.00012
					} else {
						return 0.41818
					}
				} else {
					if s[1] <= 4.0000 {
						return 0.00000
					} else {
						return 0.40000
					}
				}
			} else {
				if s[7] <= 0.6833 {
					if s[1] <= 7.5000 {
						return 0.03810
					} else {
						return 0.18571
					}
				} else {
					if s[4] <= -1.6766 {
						return 0.87500
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[0] <= 0.4719 {
				if s[4] <= -0.9828 {
					if s[2] <= 3.5000 {
						return 0.11290
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= -0.6246 {
						return 0.72414
					} else {
						return 0.38462
					}
				}
			} else {
				if s[4] <= -2.3035 {
					return 0.00000
				} else {
					if s[1] <= 4.0000 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.1953 {
			if s[2] <= 0.5000 {
				if s[3] <= -0.3820 {
					if s[8] <= 0.5000 {
						return 0.88235
					} else {
						return 0.03279
					}
				} else {
					if s[4] <= 2.7731 {
						return 0.40000
					} else {
						return 0.92453
					}
				}
			} else {
				if s[3] <= 9.2629 {
					if s[0] <= 0.2094 {
						return 0.92256
					} else {
						return 0.55556
					}
				} else {
					if s[4] <= 8.9998 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[5] <= 1.8466 {
				if s[6] <= 2.4708 {
					if s[6] <= 1.8492 {
						return 0.93827
					} else {
						return 0.74603
					}
				} else {
					if s[7] <= 0.0556 {
						return 0.98214
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 0.5000 {
					if s[6] <= 2.1239 {
						return 0.66667
					} else {
						return 0.91429
					}
				} else {
					if s[3] <= -0.3097 {
						return 0.83333
					} else {
						return 0.99539
					}
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[6] <= -0.3259 {
		if s[1] <= 23.5000 {
			if s[4] <= -1.1462 {
				if s[3] <= -3.1329 {
					if s[0] <= 0.0064 {
						return 0.00682
					} else {
						return 0.00000
					}
				} else {
					if s[8] <= 0.5000 {
						return 0.27273
					} else {
						return 0.01751
					}
				}
			} else {
				if s[7] <= 0.1944 {
					if s[5] <= -0.2693 {
						return 0.19797
					} else {
						return 0.65258
					}
				} else {
					if s[7] <= 0.7071 {
						return 0.95000
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[0] <= 0.0173 {
				if s[3] <= 0.8481 {
					if s[4] <= 3.3739 {
						return 0.02703
					} else {
						return 1.00000
					}
				} else {
					if s[1] <= 24.5000 {
						return 0.75000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[0] <= 0.1828 {
					if s[3] <= -0.6338 {
						return 0.00000
					} else {
						return 0.66667
					}
				} else {
					return 0.00000
				}
			}
		}
	} else {
		if s[3] <= -0.3159 {
			if s[4] <= 0.8868 {
				if s[4] <= -0.7887 {
					return 0.00000
				} else {
					if s[4] <= -0.4415 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 4.0000 {
					if s[4] <= 1.7083 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					if s[4] <= 1.9133 {
						return 0.66667
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[0] <= 0.5708 {
				if s[4] <= -0.0670 {
					if s[5] <= -1.4960 {
						return 0.40000
					} else {
						return 0.67308
					}
				} else {
					if s[6] <= 2.1934 {
						return 0.92111
					} else {
						return 0.99313
					}
				}
			} else {
				if s[6] <= 0.1119 {
					return 1.00000
				} else {
					if s[6] <= 0.7996 {
						return 0.20000
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[5] <= -0.3197 {
		if s[4] <= 0.0749 {
			if s[7] <= 0.5278 {
				if s[5] <= -1.1732 {
					if s[4] <= -1.6369 {
						return 0.00048
					} else {
						return 0.10526
					}
				} else {
					if s[4] <= -0.9254 {
						return 0.04545
					} else {
						return 0.46154
					}
				}
			} else {
				if s[3] <= -7.8343 {
					if s[6] <= -10.4313 {
						return 1.00000
					} else {
						return 0.09524
					}
				} else {
					if s[8] <= 0.5000 {
						return 0.81176
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[3] <= 2.4484 {
				if s[1] <= 4.5000 {
					if s[6] <= 0.7545 {
						return 0.10938
					} else {
						return 0.90909
					}
				} else {
					if s[1] <= 30.5000 {
						return 0.90741
					} else {
						return 0.69565
					}
				}
			} else {
				if s[0] <= 0.1994 {
					if s[1] <= 118.5000 {
						return 0.98500
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= -0.5515 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.4039 {
			if s[4] <= 0.0957 {
				if s[4] <= -0.9887 {
					if s[4] <= -1.9570 {
						return 0.02439
					} else {
						return 0.21212
					}
				} else {
					if s[3] <= -0.3521 {
						return 0.07143
					} else {
						return 0.60938
					}
				}
			} else {
				if s[3] <= -0.3717 {
					if s[4] <= 2.8786 {
						return 0.23077
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= 2.7462 {
						return 0.90988
					} else {
						return 0.99294
					}
				}
			}
		} else {
			if s[4] <= -0.2872 {
				if s[0] <= 0.8960 {
					if s[0] <= 0.8516 {
						return 0.04762
					} else {
						return 1.00000
					}
				} else {
					return 0.00000
				}
			} else {
				if s[0] <= 0.4280 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[4] <= -0.3221 {
		if s[3] <= -1.5563 {
			if s[7] <= 0.5278 {
				if s[4] <= -4.0405 {
					if s[0] <= 0.0064 {
						return 0.00060
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= -11.5147 {
						return 0.50000
					} else {
						return 0.00476
					}
				}
			} else {
				if s[0] <= 0.5000 {
					if s[1] <= 14.5000 {
						return 0.94915
					} else {
						return 0.00000
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[7] <= 0.2111 {
				if s[1] <= 1.5000 {
					if s[0] <= 0.1323 {
						return 0.21429
					} else {
						return 0.00746
					}
				} else {
					if s[4] <= -1.3776 {
						return 0.08966
					} else {
						return 0.43820
					}
				}
			} else {
				if s[4] <= -0.8691 {
					if s[1] <= 6.5000 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					return 1.00000
				}
			}
		}
	} else {
		if s[3] <= -0.3820 {
			if s[0] <= 0.0020 {
				if s[8] <= 0.5000 {
					return 1.00000
				} else {
					if s[3] <= -1.6602 {
						return 0.00000
					} else {
						return 0.88889
					}
				}
			} else {
				return 0.00000
			}
		} else {
			if s[1] <= 10.5000 {
				if s[5] <= 0.0207 {
					if s[6] <= -0.3356 {
						return 0.51042
					} else {
						return 0.91525
					}
				} else {
					if s[6] <= 0.3989 {
						return 0.66364
					} else {
						return 0.96119
					}
				}
			} else {
				if s[4] <= 0.7371 {
					if s[2] <= 2.5000 {
						return 0.96667
					} else {
						return 0.37500
					}
				} else {
					if s[5] <= -1.9185 {
						return 0.97959
					} else {
						return 0.99811
					}
				}
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[4] <= -0.5334 {
		if s[7] <= 0.7208 {
			if s[3] <= -1.1440 {
				if s[4] <= -3.3254 {
					if s[5] <= -7.7420 {
						return 0.00004
					} else {
						return 0.00098
					}
				} else {
					if s[3] <= -6.6074 {
						return 0.21429
					} else {
						return 0.02089
					}
				}
			} else {
				if s[4] <= -1.6320 {
					if s[0] <= 0.8655 {
						return 0.08609
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= -1.2892 {
						return 0.16364
					} else {
						return 0.63158
					}
				}
			}
		} else {
			if s[8] <= 0.5000 {
				if s[0] <= 0.5000 {
					if s[4] <= -2.5196 {
						return 1.00000
					} else {
						return 0.93333
					}
				} else {
					return 0.00000
				}
			} else {
				return 0.00000
			}
		}
	} else {
		if s[3] <= -0.3947 {
			if s[1] <= 4.5000 {
				if s[3] <= -1.5528 {
					if s[1] <= 1.5000 {
						return 0.09091
					} else {
						return 0.00000
					}
				} else {
					return 0.00000
				}
			} else {
				if s[1] <= 30.0000 {
					if s[7] <= 0.2308 {
						return 0.54545
					} else {
						return 1.00000
					}
				} else {
					if s[3] <= -1.6182 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[4] <= 1.2082 {
				if s[4] <= -0.0692 {
					if s[3] <= 3.3945 {
						return 0.60000
					} else {
						return 0.12500
					}
				} else {
					if s[4] <= 1.2002 {
						return 0.79668
					} else {
						return 0.00000
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[6] <= -3.2918 {
						return 0.57353
					} else {
						return 0.96531
					}
				} else {
					if s[4] <= 3.0759 {
						return 0.93878
					} else {
						return 0.99820
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[6] <= -0.5199 {
		if s[4] <= -0.4941 {
			if s[7] <= 0.5903 {
				if s[6] <= -0.7673 {
					if s[5] <= -0.3150 {
						return 0.00095
					} else {
						return 0.09302
					}
				} else {
					if s[5] <= -1.4940 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= -4.3450 {
					return 0.00000
				} else {
					if s[8] <= 0.5000 {
						return 0.91228
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[1] <= 4.5000 {
					if s[5] <= 2.8776 {
						return 0.10317
					} else {
						return 0.68000
					}
				} else {
					if s[5] <= 6.6087 {
						return 0.86792
					} else {
						return 0.00000
					}
				}
			} else {
				if s[4] <= 1.8552 {
					if s[5] <= 4.7344 {
						return 0.71014
					} else {
						return 0.00000
					}
				} else {
					if s[1] <= 122.5000 {
						return 0.98347
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 1.1480 {
			if s[1] <= 11.5000 {
				if s[8] <= 0.5000 {
					if s[7] <= 0.4375 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= -0.5639 {
						return 0.07895
					} else {
						return 0.72642
					}
				}
			} else {
				if s[3] <= -1.4803 {
					return 0.00000
				} else {
					if s[0] <= 0.2286 {
						return 0.98864
					} else {
						return 0.71429
					}
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[4] <= -0.3207 {
					if s[1] <= 2.5000 {
						return 0.24242
					} else {
						return 0.09091
					}
				} else {
					if s[0] <= 0.3401 {
						return 0.96939
					} else {
						return 0.33333
					}
				}
			} else {
				if s[0] <= 0.7851 {
					if s[0] <= 0.1991 {
						return 0.99723
					} else {
						return 0.91379
					}
				} else {
					return 0.00000
				}
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[6] <= -0.6143 {
		if s[4] <= 0.0701 {
			if s[4] <= -1.4301 {
				if s[1] <= 4.5000 {
					if s[7] <= 0.5000 {
						return 0.00019
					} else {
						return 0.66667
					}
				} else {
					if s[7] <= 0.7208 {
						return 0.00426
					} else {
						return 0.72917
					}
				}
			} else {
				if s[7] <= 0.1944 {
					if s[5] <= 0.6110 {
						return 0.11585
					} else {
						return 0.48485
					}
				} else {
					if s[7] <= 0.7136 {
						return 0.76923
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[1] <= 5.5000 {
				if s[4] <= 2.7743 {
					if s[5] <= -0.7722 {
						return 0.09677
					} else {
						return 0.32500
					}
				} else {
					if s[0] <= 0.0042 {
						return 0.71429
					} else {
						return 1.00000
					}
				}
			} else {
				if s[3] <= 1.7415 {
					if s[4] <= 2.0333 {
						return 0.78049
					} else {
						return 0.96875
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.80000
					} else {
						return 0.95781
					}
				}
			}
		}
	} else {
		if s[0] <= 0.5708 {
			if s[5] <= 0.9372 {
				if s[4] <= -0.3185 {
					if s[3] <= -0.4605 {
						return 0.02128
					} else {
						return 0.35714
					}
				} else {
					if s[1] <= 6.5000 {
						return 0.73333
					} else {
						return 0.96983
					}
				}
			} else {
				if s[4] <= 0.0939 {
					if s[3] <= 0.2322 {
						return 0.00000
					} else {
						return 0.71698
					}
				} else {
					if s[4] <= 2.4278 {
						return 0.94074
					} else {
						return 0.99549
					}
				}
			}
		} else {
			if s[3] <= 4.3146 {
				return 0.00000
			} else {
				return 1.00000
			}
		}
	}
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[3] <= -0.3947 {
		if s[6] <= 2.8786 {
			if s[5] <= 3.0338 {
				if s[7] <= 0.4530 {
					if s[3] <= -1.5340 {
						return 0.00032
					} else {
						return 0.09924
					}
				} else {
					if s[4] <= -4.0918 {
						return 0.00000
					} else {
						return 0.92632
					}
				}
			} else {
				if s[5] <= 3.8322 {
					return 1.00000
				} else {
					return 0.00000
				}
			}
		} else {
			if s[1] <= 5.5000 {
				return 0.00000
			} else {
				return 1.00000
			}
		}
	} else {
		if s[0] <= 0.2829 {
			if s[2] <= 0.5000 {
				if s[6] <= -0.3500 {
					if s[7] <= 0.2111 {
						return 0.32432
					} else {
						return 0.93333
					}
				} else {
					if s[5] <= -0.3753 {
						return 0.70000
					} else {
						return 0.95877
					}
				}
			} else {
				if s[4] <= -0.9284 {
					if s[3] <= 0.1626 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= 2.4329 {
						return 0.95769
					} else {
						return 0.99232
					}
				}
			}
		} else {
			if s[4] <= 0.0136 {
				if s[4] <= -0.8832 {
					if s[4] <= -2.1126 {
						return 0.01149
					} else {
						return 0.07692
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.07143
					} else {
						return 0.66667
					}
				}
			} else {
				if s[7] <= 0.1500 {
					if s[6] <= -0.0776 {
						return 0.16667
					} else {
						return 0.98182
					}
				} else {
					return 0.00000
				}
			}
		}
	}
}
