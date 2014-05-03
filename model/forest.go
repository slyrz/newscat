package model

func decisionTreeA(s boostFeature) float32 {
	if s[6] <= -0.5804 {
		if s[5] <= -2.1302 {
			if s[9] <= 0.5903 {
				if s[6] <= -3.3463 {
					if s[6] <= -4.1660 {
						return 0.00000
					} else {
						if s[2] <= 0.5000 {
							return 0.00000
						} else {
							return 0.01538
						}
					}
				} else {
					if s[6] <= -3.3038 {
						if s[7] <= -8.8044 {
							return 0.50000
						} else {
							return 0.00000
						}
					} else {
						if s[9] <= 0.2250 {
							return 0.00350
						} else {
							return 0.25000
						}
					}
				}
			} else {
				if s[6] <= -6.7818 {
					if s[4] <= 0.5000 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.5000 {
						if s[6] <= -2.3503 {
							return 1.00000
						} else {
							return 0.66667
						}
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[0] <= 0.2059 {
				if s[9] <= 0.4375 {
					if s[5] <= -0.2849 {
						if s[2] <= 0.5000 {
							return 0.02655
						} else {
							return 0.17500
						}
					} else {
						if s[6] <= -1.7210 {
							return 0.10714
						} else {
							return 0.63462
						}
					}
				} else {
					return 1.00000
				}
			} else {
				if s[1] <= 11.5000 {
					if s[8] <= -0.4351 {
						if s[5] <= -0.1509 {
							return 0.03333
						} else {
							return 0.00000
						}
					} else {
						if s[8] <= -0.3447 {
							return 1.00000
						} else {
							return 0.03704
						}
					}
				} else {
					if s[2] <= 1.0000 {
						if s[0] <= 0.6396 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[1] <= 7.5000 {
			if s[8] <= -0.4469 {
				if s[7] <= 1.0866 {
					if s[8] <= -0.5275 {
						if s[3] <= 0.5000 {
							return 0.10714
						} else {
							return 0.50000
						}
					} else {
						return 0.00000
					}
				} else {
					if s[8] <= -1.4467 {
						if s[5] <= 4.5598 {
							return 0.76000
						} else {
							return 0.25000
						}
					} else {
						if s[0] <= 0.1538 {
							return 0.00000
						} else {
							return 0.33333
						}
					}
				}
			} else {
				if s[8] <= 2.1666 {
					if s[7] <= 5.1866 {
						if s[7] <= 0.0726 {
							return 0.70732
						} else {
							return 0.85876
						}
					} else {
						if s[0] <= 0.0649 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[6] <= 0.4399 {
						if s[1] <= 3.5000 {
							return 1.00000
						} else {
							return 0.38462
						}
					} else {
						if s[5] <= 0.9737 {
							return 0.50000
						} else {
							return 0.98137
						}
					}
				}
			}
		} else {
			if s[0] <= 0.3462 {
				if s[8] <= 1.4538 {
					if s[4] <= 0.5000 {
						if s[8] <= 1.4487 {
							return 0.93735
						} else {
							return 0.00000
						}
					} else {
						if s[8] <= -0.6943 {
							return 0.88571
						} else {
							return 0.38462
						}
					}
				} else {
					if s[6] <= 0.8013 {
						if s[8] <= 4.2240 {
							return 0.85714
						} else {
							return 0.00000
						}
					} else {
						if s[8] <= 2.7508 {
							return 0.98182
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[3] <= 0.5000 {
					if s[6] <= 1.4546 {
						return 1.00000
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

func decisionTreeB(s boostFeature) float32 {
	if s[8] <= -0.5096 {
		if s[1] <= 26.5000 {
			if s[9] <= 0.5778 {
				if s[7] <= 1.0866 {
					if s[6] <= -0.5096 {
						if s[5] <= -1.6086 {
							return 0.00034
						} else {
							return 0.12054
						}
					} else {
						if s[0] <= 0.0158 {
							return 0.44660
						} else {
							return 0.13699
						}
					}
				} else {
					if s[6] <= 0.7050 {
						if s[5] <= 1.2280 {
							return 0.36364
						} else {
							return 0.01515
						}
					} else {
						if s[1] <= 6.5000 {
							return 0.68421
						} else {
							return 0.93651
						}
					}
				}
			} else {
				if s[9] <= 0.6833 {
					if s[8] <= -8.5804 {
						if s[6] <= -5.9236 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= -6.7909 {
						if s[4] <= 0.5000 {
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
			if s[7] <= -0.4249 {
				if s[5] <= -0.6738 {
					if s[3] <= 0.5000 {
						return 0.00000
					} else {
						if s[6] <= -1.7576 {
							return 0.20000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[2] <= 2.5000 {
						return 1.00000
					} else {
						if s[0] <= 0.0370 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[4] <= 0.5000 {
					return 1.00000
				} else {
					if s[2] <= 2.5000 {
						return 1.00000
					} else {
						if s[6] <= -0.1905 {
							return 0.00000
						} else {
							return 0.90909
						}
					}
				}
			}
		}
	} else {
		if s[6] <= -0.6618 {
			if s[0] <= 0.0421 {
				if s[4] <= 0.5000 {
					if s[6] <= -1.7476 {
						if s[0] <= 0.0042 {
							return 0.12500
						} else {
							return 1.00000
						}
					} else {
						if s[7] <= -2.1224 {
							return 0.75000
						} else {
							return 1.00000
						}
					}
				} else {
					return 0.00000
				}
			} else {
				if s[3] <= 0.5000 {
					if s[7] <= 3.5148 {
						return 0.00000
					} else {
						if s[1] <= 8.0000 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[0] <= 0.8333 {
						if s[5] <= -1.8160 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[0] <= 0.3512 {
				if s[6] <= 2.2212 {
					if s[8] <= 8.3206 {
						if s[0] <= 0.2118 {
							return 0.92757
						} else {
							return 0.80645
						}
					} else {
						return 0.00000
					}
				} else {
					if s[2] <= 0.5000 {
						if s[0] <= 0.0230 {
							return 0.93833
						} else {
							return 0.98718
						}
					} else {
						if s[1] <= 0.5000 {
							return 0.93750
						} else {
							return 0.99740
						}
					}
				}
			} else {
				if s[8] <= 0.3791 {
					if s[6] <= 0.0243 {
						if s[8] <= 0.1735 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						return 0.00000
					}
				} else {
					if s[7] <= 1.7649 {
						if s[1] <= 2.5000 {
							return 0.50000
						} else {
							return 0.95238
						}
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeC(s boostFeature) float32 {
	if s[7] <= -0.3658 {
		if s[6] <= -0.6680 {
			if s[6] <= -3.0571 {
				if s[1] <= 5.5000 {
					if s[3] <= 0.5000 {
						return 0.00000
					} else {
						if s[5] <= -5.5763 {
							return 0.00000
						} else {
							return 0.01556
						}
					}
				} else {
					if s[7] <= -10.0037 {
						return 0.00000
					} else {
						if s[6] <= -4.1567 {
							return 0.00532
						} else {
							return 0.14737
						}
					}
				}
			} else {
				if s[9] <= 0.4375 {
					if s[3] <= 0.5000 {
						if s[8] <= 1.3213 {
							return 0.00416
						} else {
							return 0.75000
						}
					} else {
						if s[9] <= 0.2154 {
							return 0.09392
						} else {
							return 1.00000
						}
					}
				} else {
					if s[5] <= -3.7886 {
						if s[5] <= -4.6591 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						if s[9] <= 0.9583 {
							return 1.00000
						} else {
							return 0.87500
						}
					}
				}
			}
		} else {
			if s[1] <= 5.5000 {
				if s[8] <= -0.5418 {
					if s[3] <= 0.5000 {
						if s[5] <= -0.0742 {
							return 0.06977
						} else {
							return 0.00000
						}
					} else {
						if s[5] <= -2.4382 {
							return 0.45455
						} else {
							return 0.10000
						}
					}
				} else {
					if s[7] <= -0.4933 {
						if s[8] <= 1.1385 {
							return 0.80000
						} else {
							return 0.96667
						}
					} else {
						if s[4] <= 0.5000 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[0] <= 0.0179 {
					if s[1] <= 88.0000 {
						if s[6] <= 2.7265 {
							return 0.95536
						} else {
							return 1.00000
						}
					} else {
						if s[6] <= 3.4951 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[6] <= 2.7268 {
						if s[5] <= 0.3000 {
							return 0.00000
						} else {
							return 0.66667
						}
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[6] <= 0.0709 {
			if s[8] <= -0.4351 {
				if s[6] <= -0.2016 {
					if s[7] <= 5.2704 {
						if s[1] <= 10.5000 {
							return 0.00000
						} else {
							return 0.25000
						}
					} else {
						if s[5] <= 0.8779 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[6] <= -0.0079 {
						if s[7] <= 2.3663 {
							return 0.90909
						} else {
							return 0.00000
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[5] <= 3.5861 {
					if s[0] <= 0.2063 {
						if s[3] <= 0.5000 {
							return 0.51724
						} else {
							return 0.92857
						}
					} else {
						if s[7] <= -0.0968 {
							return 1.00000
						} else {
							return 0.10345
						}
					}
				} else {
					if s[3] <= 0.5000 {
						return 0.00000
					} else {
						if s[1] <= 5.5000 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[8] <= 0.4178 {
				if s[2] <= 0.5000 {
					if s[1] <= 9.5000 {
						if s[0] <= 0.1980 {
							return 0.55882
						} else {
							return 0.14286
						}
					} else {
						if s[1] <= 13.5000 {
							return 0.88889
						} else {
							return 1.00000
						}
					}
				} else {
					if s[0] <= 0.3978 {
						if s[7] <= 1.8131 {
							return 0.82927
						} else {
							return 0.94776
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[6] <= 0.6980 {
					if s[8] <= 6.6025 {
						if s[3] <= 0.5000 {
							return 0.78571
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				} else {
					if s[2] <= 0.5000 {
						if s[9] <= 0.2125 {
							return 0.97944
						} else {
							return 0.00000
						}
					} else {
						if s[7] <= 2.2221 {
							return 0.97396
						} else {
							return 0.99884
						}
					}
				}
			}
		}
	}
}

func decisionTreeD(s boostFeature) float32 {
	if s[8] <= -0.1029 {
		if s[6] <= 0.8930 {
			if s[8] <= -0.5431 {
				if s[7] <= -0.5096 {
					if s[9] <= 0.4808 {
						if s[5] <= -1.6152 {
							return 0.00056
						} else {
							return 0.07663
						}
					} else {
						if s[9] <= 0.7136 {
							return 0.62500
						} else {
							return 0.94828
						}
					}
				} else {
					if s[1] <= 20.5000 {
						if s[6] <= 0.7050 {
							return 0.05348
						} else {
							return 1.00000
						}
					} else {
						if s[5] <= 3.2327 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[6] <= -1.5632 {
						if s[5] <= -0.9296 {
							return 0.00000
						} else {
							return 0.12500
						}
					} else {
						if s[8] <= -0.3087 {
							return 0.70000
						} else {
							return 0.00000
						}
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[2] <= 0.5000 {
				if s[0] <= 0.2171 {
					if s[1] <= 4.5000 {
						if s[6] <= 3.3373 {
							return 0.26829
						} else {
							return 1.00000
						}
					} else {
						if s[6] <= 1.0033 {
							return 0.66667
						} else {
							return 1.00000
						}
					}
				} else {
					if s[1] <= 10.0000 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= 1.5736 {
					if s[6] <= 1.4308 {
						return 1.00000
					} else {
						if s[2] <= 2.0000 {
							return 0.33333
						} else {
							return 0.00000
						}
					}
				} else {
					if s[8] <= -0.5539 {
						if s[5] <= 1.3275 {
							return 0.88889
						} else {
							return 0.97826
						}
					} else {
						if s[8] <= -0.3242 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			}
		}
	} else {
		if s[7] <= 0.4422 {
			if s[0] <= 0.2063 {
				if s[2] <= 0.5000 {
					if s[7] <= -1.7427 {
						if s[6] <= -1.7112 {
							return 0.00000
						} else {
							return 0.67391
						}
					} else {
						if s[5] <= 4.5557 {
							return 0.92500
						} else {
							return 0.00000
						}
					}
				} else {
					if s[0] <= 0.0196 {
						if s[6] <= 2.5296 {
							return 0.85714
						} else {
							return 0.98601
						}
					} else {
						if s[8] <= 1.1816 {
							return 0.36364
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[8] <= 1.9797 {
					if s[3] <= 0.5000 {
						return 0.00000
					} else {
						if s[5] <= 1.9283 {
							return 0.00000
						} else {
							return 0.75000
						}
					}
				} else {
					if s[6] <= 0.0237 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[6] <= 0.0463 {
				if s[0] <= 0.1820 {
					if s[8] <= 3.3209 {
						if s[5] <= -0.2422 {
							return 0.00000
						} else {
							return 0.90476
						}
					} else {
						if s[0] <= 0.0907 {
							return 0.22222
						} else {
							return 1.00000
						}
					}
				} else {
					if s[7] <= 1.2960 {
						if s[7] <= 1.1766 {
							return 0.16667
						} else {
							return 1.00000
						}
					} else {
						if s[6] <= -1.8648 {
							return 0.08333
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[9] <= 0.2125 {
					if s[6] <= 0.7061 {
						if s[0] <= 0.2335 {
							return 0.89474
						} else {
							return 0.45000
						}
					} else {
						if s[0] <= 0.1991 {
							return 0.99548
						} else {
							return 0.91603
						}
					}
				} else {
					if s[8] <= 2.4498 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeE(s boostFeature) float32 {
	if s[5] <= -0.2691 {
		if s[6] <= 1.0025 {
			if s[9] <= 0.4722 {
				if s[8] <= 1.4989 {
					if s[7] <= 2.4280 {
						if s[5] <= -1.6152 {
							return 0.00047
						} else {
							return 0.07639
						}
					} else {
						if s[6] <= -1.6105 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[1] <= 4.5000 {
						return 1.00000
					} else {
						if s[3] <= 0.5000 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[0] <= 0.5455 {
					if s[3] <= 0.5000 {
						if s[6] <= -1.7071 {
							return 0.61538
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= -7.1659 {
						if s[9] <= 0.9583 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[4] <= 0.5000 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[1] <= 4.0000 {
				if s[6] <= 1.1574 {
					return 1.00000
				} else {
					return 0.00000
				}
			} else {
				if s[5] <= -1.3775 {
					return 1.00000
				} else {
					if s[1] <= 22.5000 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[5] <= 1.8714 {
			if s[1] <= 5.5000 {
				if s[8] <= -0.5273 {
					if s[6] <= 2.6180 {
						if s[4] <= 0.5000 {
							return 0.03817
						} else {
							return 0.33333
						}
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.5143 {
						if s[0] <= 0.0940 {
							return 0.88462
						} else {
							return 0.57143
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[8] <= 1.7900 {
					if s[6] <= -1.5492 {
						if s[6] <= -2.2321 {
							return 0.50000
						} else {
							return 0.00000
						}
					} else {
						if s[8] <= 0.6785 {
							return 0.91538
						} else {
							return 0.70833
						}
					}
				} else {
					if s[1] <= 10.5000 {
						if s[0] <= 0.5361 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						if s[0] <= 0.1820 {
							return 1.00000
						} else {
							return 0.83333
						}
					}
				}
			}
		} else {
			if s[0] <= 0.3512 {
				if s[7] <= 2.3900 {
					if s[0] <= 0.2111 {
						if s[6] <= 0.9827 {
							return 0.49333
						} else {
							return 0.96947
						}
					} else {
						if s[8] <= -0.2830 {
							return 0.05882
						} else {
							return 0.89552
						}
					}
				} else {
					if s[2] <= 1.5000 {
						if s[6] <= 0.3912 {
							return 0.52941
						} else {
							return 0.97604
						}
					} else {
						if s[0] <= 0.2097 {
							return 0.99888
						} else {
							return 0.72727
						}
					}
				}
			} else {
				if s[7] <= -1.0156 {
					return 0.00000
				} else {
					if s[8] <= -0.4747 {
						return 0.00000
					} else {
						if s[1] <= 1.5000 {
							return 0.27778
						} else {
							return 0.51429
						}
					}
				}
			}
		}
	}
}

func decisionTreeF(s boostFeature) float32 {
	if s[8] <= -0.5418 {
		if s[1] <= 23.5000 {
			if s[7] <= 1.0820 {
				if s[9] <= 0.4688 {
					if s[6] <= -1.4905 {
						if s[5] <= -0.3067 {
							return 0.00034
						} else {
							return 0.02857
						}
					} else {
						if s[7] <= -0.8107 {
							return 0.15385
						} else {
							return 0.30159
						}
					}
				} else {
					if s[0] <= 0.5455 {
						if s[3] <= 0.5000 {
							return 0.86842
						} else {
							return 1.00000
						}
					} else {
						if s[9] <= 0.6111 {
							return 0.00000
						} else {
							return 0.63636
						}
					}
				}
			} else {
				if s[0] <= 0.2320 {
					if s[6] <= 0.7021 {
						if s[5] <= 0.6759 {
							return 0.55556
						} else {
							return 0.02500
						}
					} else {
						if s[5] <= -0.7454 {
							return 0.00000
						} else {
							return 0.85526
						}
					}
				} else {
					return 0.00000
				}
			}
		} else {
			if s[1] <= 31.5000 {
				if s[8] <= -5.0934 {
					if s[5] <= -1.7202 {
						return 0.00000
					} else {
						if s[0] <= 0.5364 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[7] <= -4.4219 {
						if s[6] <= 0.8950 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[2] <= 10.0000 {
					if s[6] <= -0.1905 {
						if s[4] <= 0.5000 {
							return 0.08333
						} else {
							return 0.33333
						}
					} else {
						if s[2] <= 2.5000 {
							return 1.00000
						} else {
							return 0.94595
						}
					}
				} else {
					if s[7] <= 3.2858 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[5] <= -0.4871 {
			if s[8] <= 1.8775 {
				return 0.00000
			} else {
				return 1.00000
			}
		} else {
			if s[0] <= 0.4472 {
				if s[2] <= 0.5000 {
					if s[3] <= 0.5000 {
						if s[6] <= 0.5656 {
							return 0.48750
						} else {
							return 0.95752
						}
					} else {
						if s[0] <= 0.2130 {
							return 0.97658
						} else {
							return 0.82540
						}
					}
				} else {
					if s[6] <= 0.4184 {
						if s[0] <= 0.2063 {
							return 0.75000
						} else {
							return 0.00000
						}
					} else {
						if s[5] <= 0.0764 {
							return 0.71429
						} else {
							return 0.99334
						}
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[5] <= 3.1143 {
						if s[5] <= 3.0882 {
							return 0.04545
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				} else {
					if s[1] <= 8.5000 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeG(s boostFeature) float32 {
	if s[5] <= -0.4279 {
		if s[5] <= -3.4517 {
			if s[5] <= -5.5368 {
				if s[6] <= -0.8836 {
					if s[8] <= -2.8970 {
						if s[9] <= 0.8611 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[3] <= 0.5000 {
							return 0.00000
						} else {
							return 0.18182
						}
					}
				} else {
					if s[6] <= -0.2939 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[3] <= 0.5000 {
					if s[1] <= 4.5000 {
						return 0.00000
					} else {
						if s[9] <= 0.4182 {
							return 0.01674
						} else {
							return 1.00000
						}
					}
				} else {
					if s[9] <= 0.4167 {
						if s[2] <= 0.5000 {
							return 0.00476
						} else {
							return 0.08108
						}
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[5] <= -1.2174 {
				if s[1] <= 4.5000 {
					if s[6] <= 0.9036 {
						if s[3] <= 0.5000 {
							return 0.00000
						} else {
							return 0.02419
						}
					} else {
						return 1.00000
					}
				} else {
					if s[9] <= 0.4818 {
						if s[2] <= 0.5000 {
							return 0.04478
						} else {
							return 0.12329
						}
					} else {
						if s[6] <= -2.2290 {
							return 0.84615
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[0] <= 0.0109 {
					if s[1] <= 5.0000 {
						return 0.00000
					} else {
						if s[9] <= 0.5625 {
							return 0.61538
						} else {
							return 1.00000
						}
					}
				} else {
					return 0.00000
				}
			}
		}
	} else {
		if s[1] <= 9.5000 {
			if s[6] <= 0.4591 {
				if s[0] <= 0.2059 {
					if s[5] <= 0.2095 {
						if s[8] <= -6.8806 {
							return 0.68750
						} else {
							return 1.00000
						}
					} else {
						if s[0] <= 0.1159 {
							return 0.33333
						} else {
							return 0.73333
						}
					}
				} else {
					if s[8] <= -0.5077 {
						if s[7] <= -4.1177 {
							return 0.18750
						} else {
							return 0.00704
						}
					} else {
						if s[8] <= 4.7481 {
							return 0.32787
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[8] <= 0.4134 {
					if s[6] <= 3.3101 {
						if s[9] <= 0.1929 {
							return 0.46218
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				} else {
					if s[1] <= 7.5000 {
						if s[6] <= 1.0469 {
							return 0.88525
						} else {
							return 0.97608
						}
					} else {
						if s[0] <= 0.5841 {
							return 0.97959
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[7] <= 0.8132 {
				if s[6] <= -0.3535 {
					if s[6] <= -1.0196 {
						if s[6] <= -4.0544 {
							return 0.20000
						} else {
							return 0.00000
						}
					} else {
						if s[7] <= -2.5325 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[1] <= 10.5000 {
						if s[5] <= 3.1709 {
							return 0.57143
						} else {
							return 1.00000
						}
					} else {
						if s[5] <= 1.5616 {
							return 0.84783
						} else {
							return 0.99507
						}
					}
				}
			} else {
				if s[7] <= 4.4448 {
					if s[5] <= -0.2577 {
						if s[6] <= -1.4859 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[7] <= 4.4431 {
							return 0.98294
						} else {
							return 0.00000
						}
					}
				} else {
					if s[2] <= 0.5000 {
						if s[6] <= 1.3780 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[2] <= 3.5000 {
							return 0.99725
						} else {
							return 1.00000
						}
					}
				}
			}
		}
	}
}

func decisionTreeH(s boostFeature) float32 {
	if s[5] <= -0.2691 {
		if s[0] <= 0.1090 {
			if s[1] <= 4.5000 {
				if s[6] <= 3.3023 {
					if s[9] <= 0.4375 {
						if s[0] <= 0.1082 {
							return 0.00108
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				} else {
					return 1.00000
				}
			} else {
				if s[6] <= -1.2056 {
					if s[9] <= 0.5433 {
						if s[5] <= -1.6086 {
							return 0.00348
						} else {
							return 0.34375
						}
					} else {
						return 1.00000
					}
				} else {
					if s[1] <= 16.0000 {
						if s[7] <= -1.5245 {
							return 0.79630
						} else {
							return 0.29412
						}
					} else {
						if s[5] <= -1.0146 {
							return 0.00000
						} else {
							return 0.40000
						}
					}
				}
			}
		} else {
			if s[5] <= -0.3786 {
				if s[8] <= -8.9420 {
					if s[9] <= 0.8091 {
						return 0.00000
					} else {
						if s[6] <= -6.6053 {
							return 0.25000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[9] <= 0.4583 {
						return 0.00000
					} else {
						if s[6] <= -3.5422 {
							return 1.00000
						} else {
							return 0.33333
						}
					}
				}
			} else {
				if s[6] <= -1.4961 {
					return 0.00000
				} else {
					return 1.00000
				}
			}
		}
	} else {
		if s[6] <= 0.4267 {
			if s[6] <= -1.3243 {
				if s[0] <= 0.0405 {
					if s[1] <= 5.5000 {
						if s[7] <= 3.4622 {
							return 0.00000
						} else {
							return 0.33333
						}
					} else {
						if s[6] <= -2.2182 {
							return 0.81818
						} else {
							return 0.00000
						}
					}
				} else {
					if s[0] <= 0.4665 {
						if s[0] <= 0.4547 {
							return 0.03333
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[5] <= 3.0683 {
					if s[8] <= -0.5077 {
						if s[5] <= 1.1927 {
							return 0.62712
						} else {
							return 0.00000
						}
					} else {
						if s[6] <= 0.3986 {
							return 0.86957
						} else {
							return 0.00000
						}
					}
				} else {
					if s[5] <= 5.0772 {
						if s[3] <= 0.5000 {
							return 0.00000
						} else {
							return 0.36842
						}
					} else {
						if s[5] <= 5.6996 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[8] <= 0.4178 {
				if s[1] <= 5.5000 {
					if s[6] <= 2.9885 {
						if s[0] <= 0.0087 {
							return 0.11765
						} else {
							return 0.48276
						}
					} else {
						if s[0] <= 0.0142 {
							return 0.80000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[6] <= 0.8868 {
						if s[5] <= 2.2191 {
							return 0.80000
						} else {
							return 0.16667
						}
					} else {
						if s[5] <= 3.5448 {
							return 0.98204
						} else {
							return 0.91139
						}
					}
				}
			} else {
				if s[9] <= 0.1010 {
					if s[1] <= 19.5000 {
						if s[7] <= 0.0310 {
							return 0.86957
						} else {
							return 0.97722
						}
					} else {
						if s[5] <= 0.9670 {
							return 0.93750
						} else {
							return 1.00000
						}
					}
				} else {
					if s[6] <= 5.3380 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			}
		}
	}
}

func decisionTreeI(s boostFeature) float32 {
	if s[1] <= 15.5000 {
		if s[8] <= 0.4563 {
			if s[1] <= 6.5000 {
				if s[6] <= -0.1655 {
					if s[6] <= -1.4645 {
						if s[8] <= -4.0421 {
							return 0.00029
						} else {
							return 0.01159
						}
					} else {
						if s[8] <= -0.5918 {
							return 0.04795
						} else {
							return 0.55556
						}
					}
				} else {
					if s[3] <= 0.5000 {
						if s[5] <= -0.1204 {
							return 0.05263
						} else {
							return 0.32927
						}
					} else {
						if s[9] <= 0.4841 {
							return 0.38182
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[7] <= -1.0451 {
					if s[0] <= 0.0510 {
						if s[9] <= 0.2679 {
							return 0.03396
						} else {
							return 0.97000
						}
					} else {
						if s[6] <= -5.2564 {
							return 0.00000
						} else {
							return 0.13978
						}
					}
				} else {
					if s[5] <= -0.7179 {
						return 0.00000
					} else {
						if s[8] <= 0.2262 {
							return 0.73034
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[1] <= 0.5000 {
				if s[7] <= -1.4277 {
					return 0.00000
				} else {
					if s[0] <= 0.2418 {
						if s[5] <= 0.3683 {
							return 0.00000
						} else {
							return 0.98077
						}
					} else {
						if s[3] <= 0.5000 {
							return 0.11111
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[6] <= -1.7516 {
					if s[0] <= 0.1425 {
						if s[1] <= 5.0000 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[6] <= -2.0237 {
							return 0.00000
						} else {
							return 0.33333
						}
					}
				} else {
					if s[6] <= 0.3752 {
						if s[5] <= 3.2907 {
							return 0.86842
						} else {
							return 0.30556
						}
					} else {
						if s[5] <= 0.5393 {
							return 0.37500
						} else {
							return 0.97917
						}
					}
				}
			}
		}
	} else {
		if s[1] <= 23.5000 {
			if s[7] <= -0.6119 {
				if s[6] <= 0.7740 {
					if s[3] <= 0.5000 {
						return 0.00000
					} else {
						if s[8] <= -2.9590 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[0] <= 0.0630 {
						return 1.00000
					} else {
						if s[6] <= 2.8572 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[6] <= -0.1896 {
					return 0.00000
				} else {
					if s[5] <= 1.5616 {
						if s[5] <= 1.4446 {
							return 1.00000
						} else {
							return 0.50000
						}
					} else {
						if s[8] <= 2.2560 {
							return 0.91954
						} else {
							return 1.00000
						}
					}
				}
			}
		} else {
			if s[7] <= -3.5256 {
				if s[5] <= -0.9577 {
					if s[5] <= -2.2398 {
						return 0.00000
					} else {
						if s[4] <= 0.5000 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					return 1.00000
				}
			} else {
				if s[7] <= 4.0095 {
					if s[5] <= -0.6024 {
						return 0.00000
					} else {
						if s[6] <= 0.4069 {
							return 0.50000
						} else {
							return 0.99248
						}
					}
				} else {
					if s[0] <= 0.2111 {
						return 1.00000
					} else {
						if s[0] <= 0.2263 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			}
		}
	}
}

func decisionTreeJ(s boostFeature) float32 {
	if s[6] <= -0.6855 {
		if s[5] <= -2.4245 {
			if s[0] <= 0.0064 {
				if s[3] <= 0.5000 {
					if s[8] <= -6.3510 {
						if s[7] <= -3.3809 {
							return 0.00150
						} else {
							return 0.00971
						}
					} else {
						if s[9] <= 0.4182 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[9] <= 0.4792 {
						if s[6] <= -5.4005 {
							return 0.00000
						} else {
							return 0.00885
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= -5.6619 {
					return 0.00000
				} else {
					if s[6] <= -5.5387 {
						if s[9] <= 0.4048 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[7] <= -7.6605 {
							return 0.03571
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[6] <= -1.6012 {
				if s[5] <= -2.4202 {
					return 1.00000
				} else {
					if s[1] <= 6.5000 {
						if s[6] <= -1.7510 {
							return 0.00362
						} else {
							return 0.11111
						}
					} else {
						if s[0] <= 0.4550 {
							return 0.38298
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[3] <= 0.5000 {
					if s[5] <= -1.2174 {
						return 0.00000
					} else {
						if s[0] <= 0.0109 {
							return 0.50000
						} else {
							return 0.08696
						}
					}
				} else {
					if s[0] <= 0.1373 {
						if s[2] <= 1.5000 {
							return 0.60000
						} else {
							return 0.00000
						}
					} else {
						if s[4] <= 0.5000 {
							return 0.91667
						} else {
							return 0.00000
						}
					}
				}
			}
		}
	} else {
		if s[5] <= -0.4279 {
			if s[7] <= -4.4285 {
				if s[9] <= 0.4188 {
					if s[7] <= -4.5786 {
						if s[8] <= 0.0174 {
							return 0.05714
						} else {
							return 0.50000
						}
					} else {
						return 1.00000
					}
				} else {
					return 1.00000
				}
			} else {
				if s[6] <= 3.7670 {
					if s[6] <= 1.0253 {
						return 0.00000
					} else {
						if s[9] <= 0.4500 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[0] <= 0.2111 {
				if s[8] <= -0.0224 {
					if s[7] <= -0.5292 {
						if s[5] <= 2.8212 {
							return 0.59000
						} else {
							return 0.93750
						}
					} else {
						if s[5] <= 5.5489 {
							return 0.90071
						} else {
							return 0.67188
						}
					}
				} else {
					if s[6] <= 0.0254 {
						if s[6] <= -0.1746 {
							return 0.90909
						} else {
							return 0.16667
						}
					} else {
						if s[8] <= 2.3722 {
							return 0.94954
						} else {
							return 0.99455
						}
					}
				}
			} else {
				if s[3] <= 0.5000 {
					if s[6] <= 2.6693 {
						if s[8] <= -0.4469 {
							return 0.02703
						} else {
							return 0.50602
						}
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= -0.1226 {
						if s[1] <= 2.5000 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[5] <= 7.0422 {
							return 0.97895
						} else {
							return 0.00000
						}
					}
				}
			}
		}
	}
}
