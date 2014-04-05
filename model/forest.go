package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[4] <= 0.1761 {
		if s[6] <= -0.7723 {
			if s[7] <= 0.5500 {
				if s[6] <= -2.0910 {
					if s[5] <= 0.3660 {
						return 0.00085
					} else {
						return 0.16495
					}
				} else {
					if s[4] <= -1.4730 {
						return 0.00000
					} else {
						return 0.30769
					}
				}
			} else {
				if s[0] <= 0.5455 {
					if s[4] <= -3.6186 {
						return 1.00000
					} else {
						return 0.85185
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[4] <= -1.2504 {
				if s[6] <= 0.6465 {
					return 0.00000
				} else {
					if s[6] <= 0.6671 {
						return 1.00000
					} else {
						return 0.11111
					}
				}
			} else {
				if s[1] <= 17.5000 {
					if s[6] <= 4.9026 {
						return 0.75000
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= 1.1913 {
						return 0.00000
					} else {
						return 0.25000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.3180 {
			if s[4] <= 2.1913 {
				if s[7] <= 0.0801 {
					if s[1] <= 9.5000 {
						return 0.28696
					} else {
						return 0.84314
					}
				} else {
					return 1.00000
				}
			} else {
				if s[5] <= 2.1117 {
					if s[5] <= 1.9583 {
						return 0.84314
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 8.9591 {
						return 0.96992
					} else {
						return 0.83333
					}
				}
			}
		} else {
			if s[4] <= 2.6982 {
				if s[7] <= 0.0556 {
					if s[3] <= 5.6001 {
						return 0.92133
					} else {
						return 0.44444
					}
				} else {
					return 0.00000
				}
			} else {
				if s[1] <= 2.5000 {
					if s[0] <= 0.3404 {
						return 0.96113
					} else {
						return 0.00000
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.99759
					} else {
						return 1.00000
					}
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[4] <= -0.3610 {
		if s[3] <= -1.8798 {
			if s[3] <= -4.0374 {
				if s[6] <= -2.8489 {
					if s[4] <= -2.2114 {
						return 0.00066
					} else {
						return 0.13333
					}
				} else {
					if s[6] <= -2.8157 {
						return 0.75000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[3] <= -4.0341 {
					if s[7] <= 0.4375 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= -1.5073 {
						return 0.02706
					} else {
						return 0.14286
					}
				}
			}
		} else {
			if s[4] <= -1.8756 {
				if s[1] <= 9.5000 {
					if s[4] <= -2.1067 {
						return 0.00000
					} else {
						return 0.10000
					}
				} else {
					if s[3] <= 1.1299 {
						return 0.00000
					} else {
						return 0.54545
					}
				}
			} else {
				if s[1] <= 6.5000 {
					if s[6] <= -3.1315 {
						return 0.01149
					} else {
						return 0.35385
					}
				} else {
					if s[0] <= 0.2059 {
						return 0.55714
					} else {
						return 0.05556
					}
				}
			}
		}
	} else {
		if s[6] <= 0.6495 {
			if s[1] <= 6.5000 {
				if s[0] <= 0.0087 {
					if s[5] <= 4.9027 {
						return 0.34211
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= -0.3667 {
						return 0.10811
					} else {
						return 0.80597
					}
				}
			} else {
				if s[4] <= 1.6401 {
					if s[3] <= 2.7822 {
						return 0.81373
					} else {
						return 0.36842
					}
				} else {
					if s[5] <= 4.3868 {
						return 0.90071
					} else {
						return 0.99482
					}
				}
			}
		} else {
			if s[4] <= 1.4459 {
				if s[5] <= -1.4331 {
					if s[5] <= -3.5544 {
						return 0.72727
					} else {
						return 0.00000
					}
				} else {
					if s[6] <= 6.6599 {
						return 0.91667
					} else {
						return 0.16667
					}
				}
			} else {
				if s[4] <= 3.2236 {
					if s[7] <= 0.0556 {
						return 0.96234
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 1.3242 {
						return 0.95349
					} else {
						return 0.99826
					}
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[3] <= -0.4127 {
		if s[4] <= -0.7709 {
			if s[7] <= 0.4808 {
				if s[3] <= -0.7020 {
					if s[1] <= 47.5000 {
						return 0.00085
					} else {
						return 0.30769
					}
				} else {
					if s[3] <= -0.6998 {
						return 1.00000
					} else {
						return 0.15385
					}
				}
			} else {
				if s[0] <= 0.5455 {
					if s[4] <= -3.5698 {
						return 1.00000
					} else {
						return 0.86667
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[5] <= -0.5160 {
				if s[7] <= 0.1534 {
					if s[5] <= -10.2761 {
						return 1.00000
					} else {
						return 0.05376
					}
				} else {
					if s[1] <= 8.0000 {
						return 0.50000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[3] <= -1.1977 {
					if s[0] <= 0.1393 {
						return 0.80000
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= -0.8827 {
						return 1.00000
					} else {
						return 0.69231
					}
				}
			}
		}
	} else {
		if s[0] <= 0.4073 {
			if s[6] <= 0.0099 {
				if s[2] <= 0.5000 {
					if s[4] <= 3.2386 {
						return 0.32877
					} else {
						return 0.94872
					}
				} else {
					if s[5] <= 2.4570 {
						return 0.67273
					} else {
						return 0.93833
					}
				}
			} else {
				if s[4] <= 0.3734 {
					if s[6] <= 4.9026 {
						return 0.66279
					} else {
						return 0.08333
					}
				} else {
					if s[6] <= 2.0496 {
						return 0.93734
					} else {
						return 0.99037
					}
				}
			}
		} else {
			if s[1] <= 9.5000 {
				if s[3] <= 0.8169 {
					if s[5] <= -7.3107 {
						return 1.00000
					} else {
						return 0.05000
					}
				} else {
					if s[0] <= 0.6905 {
						return 0.13514
					} else {
						return 0.00833
					}
				}
			} else {
				if s[4] <= -1.5584 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[6] <= -0.5688 {
		if s[5] <= 0.3682 {
			if s[3] <= -1.6631 {
				if s[3] <= -4.1105 {
					if s[8] <= 0.5000 {
						return 0.00237
					} else {
						return 0.00006
					}
				} else {
					if s[4] <= 0.4570 {
						return 0.03013
					} else {
						return 0.83333
					}
				}
			} else {
				if s[8] <= 0.5000 {
					if s[1] <= 3.5000 {
						return 0.00000
					} else {
						return 0.78571
					}
				} else {
					if s[4] <= 1.6500 {
						return 0.11526
					} else {
						return 0.73684
					}
				}
			}
		} else {
			if s[3] <= 0.0154 {
				if s[1] <= 11.0000 {
					if s[4] <= 0.7876 {
						return 0.00000
					} else {
						return 0.50000
					}
				} else {
					if s[2] <= 6.5000 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[7] <= 0.0714 {
						return 0.35185
					} else {
						return 1.00000
					}
				} else {
					if s[3] <= 7.4300 {
						return 0.90040
					} else {
						return 0.50000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.2341 {
			if s[4] <= -0.4717 {
				if s[3] <= 0.1889 {
					return 0.00000
				} else {
					if s[3] <= 3.1069 {
						return 0.54839
					} else {
						return 0.07143
					}
				}
			} else {
				if s[4] <= 2.6977 {
					if s[3] <= -0.4848 {
						return 0.64706
					} else {
						return 0.93770
					}
				} else {
					if s[4] <= 3.6846 {
						return 0.97985
					} else {
						return 0.99908
					}
				}
			}
		} else {
			if s[4] <= 0.3172 {
				if s[4] <= -2.1586 {
					return 0.00000
				} else {
					if s[1] <= 4.5000 {
						return 0.16667
					} else {
						return 0.75000
					}
				}
			} else {
				if s[5] <= -0.0838 {
					if s[0] <= 0.5583 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[3] <= 0.1335 {
						return 0.00000
					} else {
						return 0.88776
					}
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[5] <= -0.3476 {
		if s[3] <= -0.7631 {
			if s[7] <= 0.4530 {
				if s[1] <= 47.5000 {
					if s[4] <= -1.6409 {
						return 0.00030
					} else {
						return 0.06952
					}
				} else {
					if s[1] <= 50.0000 {
						return 0.50000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[8] <= 0.5000 {
					if s[0] <= 0.5000 {
						return 0.93220
					} else {
						return 0.00000
					}
				} else {
					if s[0] <= 0.5455 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[7] <= 0.1181 {
					if s[6] <= -1.6144 {
						return 0.05970
					} else {
						return 0.50847
					}
				} else {
					if s[3] <= -0.6592 {
						return 0.33333
					} else {
						return 1.00000
					}
				}
			} else {
				if s[1] <= 7.5000 {
					if s[4] <= -0.4528 {
						return 0.04762
					} else {
						return 0.78571
					}
				} else {
					if s[5] <= -1.6353 {
						return 0.89714
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[5] <= 2.3626 {
			if s[6] <= -1.1332 {
				if s[4] <= -0.5192 {
					if s[4] <= -1.3331 {
						return 0.00000
					} else {
						return 0.17391
					}
				} else {
					if s[5] <= 2.2656 {
						return 0.73684
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 5.5000 {
					if s[4] <= -0.4717 {
						return 0.18421
					} else {
						return 0.91450
					}
				} else {
					if s[4] <= -3.8297 {
						return 0.00000
					} else {
						return 0.98031
					}
				}
			}
		} else {
			if s[1] <= 0.5000 {
				if s[6] <= 2.9616 {
					if s[2] <= 0.5000 {
						return 0.14815
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= 4.7941 {
						return 1.00000
					} else {
						return 0.71429
					}
				}
			} else {
				if s[6] <= -0.1779 {
					if s[1] <= 8.5000 {
						return 0.48750
					} else {
						return 0.91288
					}
				} else {
					if s[0] <= 0.6985 {
						return 0.98628
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[4] <= -0.4124 {
		if s[7] <= 0.4808 {
			if s[7] <= 0.4202 {
				if s[5] <= 0.9257 {
					if s[0] <= 0.0940 {
						return 0.00434
					} else {
						return 0.00021
					}
				} else {
					if s[1] <= 1.5000 {
						return 0.00000
					} else {
						return 0.32432
					}
				}
			} else {
				if s[4] <= -6.6044 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		} else {
			if s[3] <= -9.6708 {
				return 0.00000
			} else {
				if s[7] <= 0.8516 {
					if s[1] <= 10.0000 {
						return 0.71875
					} else {
						return 0.22222
					}
				} else {
					if s[7] <= 0.9199 {
						return 1.00000
					} else {
						return 0.66667
					}
				}
			}
		}
	} else {
		if s[3] <= -0.4848 {
			if s[5] <= -0.2366 {
				if s[8] <= 0.5000 {
					if s[1] <= 2.0000 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.0516 {
						return 0.00000
					} else {
						return 0.19231
					}
				}
			} else {
				if s[0] <= 0.1876 {
					if s[6] <= -0.4964 {
						return 0.20000
					} else {
						return 1.00000
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[0] <= 0.2341 {
				if s[4] <= 0.6331 {
					if s[5] <= 4.1682 {
						return 0.79032
					} else {
						return 0.26667
					}
				} else {
					if s[2] <= 0.5000 {
						return 0.94620
					} else {
						return 0.98916
					}
				}
			} else {
				if s[0] <= 0.2352 {
					return 0.00000
				} else {
					if s[6] <= -6.0432 {
						return 0.00000
					} else {
						return 0.89362
					}
				}
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[5] <= -0.0845 {
		if s[4] <= 0.7989 {
			if s[7] <= 0.4722 {
				if s[6] <= 0.4218 {
					if s[1] <= 47.5000 {
						return 0.00155
					} else {
						return 0.15000
					}
				} else {
					if s[4] <= -0.4784 {
						return 0.16418
					} else {
						return 0.80000
					}
				}
			} else {
				if s[0] <= 0.5455 {
					if s[1] <= 10.5000 {
						return 0.96552
					} else {
						return 0.76471
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[1] <= 5.5000 {
				if s[8] <= 0.5000 {
					return 1.00000
				} else {
					if s[5] <= -0.7712 {
						return 0.21212
					} else {
						return 0.60000
					}
				}
			} else {
				if s[6] <= -9.2230 {
					if s[1] <= 60.5000 {
						return 0.89362
					} else {
						return 0.25000
					}
				} else {
					if s[1] <= 10.5000 {
						return 0.93939
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.8252 {
			if s[1] <= 8.5000 {
				if s[0] <= 0.4046 {
					if s[0] <= 0.0277 {
						return 0.26437
					} else {
						return 0.66990
					}
				} else {
					if s[4] <= 0.0895 {
						return 0.01538
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= -1.2023 {
					return 0.00000
				} else {
					if s[3] <= -2.1045 {
						return 0.00000
					} else {
						return 0.95977
					}
				}
			}
		} else {
			if s[3] <= -0.5064 {
				if s[4] <= 0.0913 {
					if s[0] <= 0.1876 {
						return 0.50000
					} else {
						return 0.00000
					}
				} else {
					return 1.00000
				}
			} else {
				if s[4] <= 0.3502 {
					if s[3] <= 3.0513 {
						return 0.75000
					} else {
						return 0.30303
					}
				} else {
					if s[0] <= 0.6807 {
						return 0.98792
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[3] <= -0.4136 {
		if s[4] <= -1.6419 {
			if s[4] <= -4.5861 {
				if s[4] <= -6.5452 {
					return 0.00000
				} else {
					if s[5] <= -9.9943 {
						return 0.00252
					} else {
						return 0.00000
					}
				}
			} else {
				if s[7] <= 0.4653 {
					if s[1] <= 48.0000 {
						return 0.00686
					} else {
						return 0.25000
					}
				} else {
					if s[1] <= 7.5000 {
						return 1.00000
					} else {
						return 0.72222
					}
				}
			}
		} else {
			if s[7] <= 0.0940 {
				if s[1] <= 9.5000 {
					if s[5] <= 0.4179 {
						return 0.03191
					} else {
						return 0.69231
					}
				} else {
					if s[4] <= 0.5724 {
						return 0.27778
					} else {
						return 0.90000
					}
				}
			} else {
				if s[3] <= -0.6125 {
					if s[6] <= -9.2230 {
						return 0.87805
					} else {
						return 1.00000
					}
				} else {
					return 0.00000
				}
			}
		}
	} else {
		if s[4] <= -0.4035 {
			if s[6] <= -2.0203 {
				if s[1] <= 15.5000 {
					if s[2] <= 0.5000 {
						return 0.00000
					} else {
						return 0.10000
					}
				} else {
					if s[0] <= 0.1056 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[3] <= 0.8169 {
					if s[6] <= -1.7401 {
						return 1.00000
					} else {
						return 0.40000
					}
				} else {
					if s[4] <= -2.3531 {
						return 0.01724
					} else {
						return 0.37931
					}
				}
			}
		} else {
			if s[6] <= -0.4834 {
				if s[1] <= 6.5000 {
					if s[4] <= 2.5013 {
						return 0.22078
					} else {
						return 0.84375
					}
				} else {
					if s[4] <= 1.7409 {
						return 0.62500
					} else {
						return 0.93462
					}
				}
			} else {
				if s[5] <= 1.6191 {
					if s[2] <= 0.5000 {
						return 0.85841
					} else {
						return 0.96875
					}
				} else {
					if s[7] <= 0.2409 {
						return 0.98615
					} else {
						return 0.40000
					}
				}
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[4] <= -0.4623 {
		if s[3] <= -1.6324 {
			if s[7] <= 0.4808 {
				if s[6] <= -0.7685 {
					if s[3] <= -1.9644 {
						return 0.00026
					} else {
						return 0.05357
					}
				} else {
					if s[6] <= -0.7269 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 9.5000 {
					if s[4] <= -6.4663 {
						return 0.00000
					} else {
						return 0.92500
					}
				} else {
					if s[0] <= 0.5000 {
						return 0.75000
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[8] <= 0.5000 {
				if s[6] <= -6.1698 {
					if s[7] <= 0.1458 {
						return 0.31250
					} else {
						return 0.95000
					}
				} else {
					if s[4] <= -0.9910 {
						return 0.00000
					} else {
						return 0.66667
					}
				}
			} else {
				if s[4] <= -2.0938 {
					if s[6] <= 6.6476 {
						return 0.00524
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= -2.0310 {
						return 0.09615
					} else {
						return 0.36986
					}
				}
			}
		}
	} else {
		if s[3] <= -0.4231 {
			if s[1] <= 5.5000 {
				if s[6] <= 0.8996 {
					if s[7] <= 0.3000 {
						return 0.03175
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= 0.1220 {
						return 0.00000
					} else {
						return 0.80000
					}
				}
			} else {
				if s[0] <= 0.0116 {
					if s[2] <= 1.5000 {
						return 0.96552
					} else {
						return 0.60000
					}
				} else {
					if s[0] <= 0.1228 {
						return 0.25000
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[4] <= 1.9673 {
				if s[7] <= 0.2361 {
					if s[5] <= -0.4012 {
						return 0.48315
					} else {
						return 0.84597
					}
				} else {
					return 1.00000
				}
			} else {
				if s[6] <= 0.4497 {
					if s[1] <= 8.5000 {
						return 0.69118
					} else {
						return 0.95203
					}
				} else {
					if s[7] <= 0.2214 {
						return 0.99413
					} else {
						return 0.66667
					}
				}
			}
		}
	}
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[4] <= 0.2303 {
		if s[4] <= -1.6419 {
			if s[4] <= -3.9264 {
				if s[7] <= 0.7071 {
					return 0.00000
				} else {
					if s[2] <= 1.0000 {
						return 0.08333
					} else {
						return 1.00000
					}
				}
			} else {
				if s[4] <= -3.9214 {
					return 1.00000
				} else {
					if s[6] <= 5.2666 {
						return 0.02372
					} else {
						return 0.66667
					}
				}
			}
		} else {
			if s[5] <= -0.4768 {
				if s[7] <= 0.1181 {
					if s[2] <= 0.5000 {
						return 0.12626
					} else {
						return 0.26087
					}
				} else {
					if s[8] <= 0.5000 {
						return 0.92105
					} else {
						return 0.00000
					}
				}
			} else {
				if s[3] <= 3.1542 {
					if s[3] <= -1.7402 {
						return 0.00000
					} else {
						return 0.64646
					}
				} else {
					if s[0] <= 0.5019 {
						return 0.09302
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[4] <= 1.6843 {
			if s[3] <= 1.0612 {
				if s[2] <= 0.5000 {
					if s[6] <= -0.2749 {
						return 0.21277
					} else {
						return 1.00000
					}
				} else {
					if s[2] <= 1.5000 {
						return 0.95652
					} else {
						return 0.42857
					}
				}
			} else {
				if s[4] <= 1.6788 {
					if s[1] <= 2.5000 {
						return 0.88889
					} else {
						return 0.79333
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[6] <= -8.1731 {
				if s[1] <= 15.5000 {
					if s[4] <= 3.1456 {
						return 0.42105
					} else {
						return 0.74359
					}
				} else {
					if s[3] <= -0.6519 {
						return 0.00000
					} else {
						return 0.97183
					}
				}
			} else {
				if s[5] <= 0.2431 {
					if s[1] <= 7.5000 {
						return 0.64286
					} else {
						return 0.97549
					}
				} else {
					if s[1] <= 0.5000 {
						return 0.87500
					} else {
						return 0.99405
					}
				}
			}
		}
	}
}
