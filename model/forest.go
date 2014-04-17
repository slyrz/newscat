package model

func decisionTreeA(s ScoreFeature) float32 {
	if s[6] <= -0.3537 {
		if s[6] <= -1.3043 {
			if s[9] <= 0.5357 {
				if s[6] <= -1.7893 {
					if s[7] <= 3.2200 {
						if s[5] <= -3.4585 {
							return 0.00022
						} else {
							return 0.01613
						}
					} else {
						if s[7] <= 3.2852 {
							return 0.50000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[8] <= 1.6237 {
						if s[0] <= 0.7865 {
							return 0.03289
						} else {
							return 1.00000
						}
					} else {
						if s[7] <= -2.5441 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[5] <= -8.8873 {
					return 0.00000
				} else {
					if s[9] <= 0.8167 {
						return 1.00000
					} else {
						if s[6] <= -6.4589 {
							return 0.00000
						} else {
							return 0.89474
						}
					}
				}
			}
		} else {
			if s[8] <= -0.7942 {
				if s[9] <= 0.2792 {
					if s[0] <= 0.7952 {
						if s[9] <= 0.0385 {
							return 0.10687
						} else {
							return 0.00000
						}
					} else {
						return 1.00000
					}
				} else {
					return 1.00000
				}
			} else {
				if s[5] <= 2.7411 {
					if s[0] <= 0.0940 {
						if s[2] <= 0.5000 {
							return 0.96154
						} else {
							return 0.33333
						}
					} else {
						if s[7] <= 0.3093 {
							return 0.50000
						} else {
							return 0.12500
						}
					}
				} else {
					if s[7] <= 1.5596 {
						return 0.00000
					} else {
						if s[7] <= 4.2497 {
							return 0.83333
						} else {
							return 0.00000
						}
					}
				}
			}
		}
	} else {
		if s[6] <= 1.2044 {
			if s[7] <= -0.9865 {
				if s[7] <= -5.2758 {
					if s[9] <= 0.2538 {
						if s[7] <= -8.6336 {
							return 0.35714
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				} else {
					if s[7] <= -1.3703 {
						if s[5] <= -0.1844 {
							return 0.00000
						} else {
							return 0.07692
						}
					} else {
						if s[7] <= -1.1773 {
							return 0.66667
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[1] <= 26.5000 {
					if s[5] <= 4.2411 {
						if s[0] <= 0.2438 {
							return 0.87500
						} else {
							return 0.66667
						}
					} else {
						if s[7] <= 0.4618 {
							return 0.00000
						} else {
							return 0.52632
						}
					}
				} else {
					if s[1] <= 71.0000 {
						if s[8] <= -5.0568 {
							return 0.00000
						} else {
							return 0.40000
						}
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[8] <= 2.2368 {
				if s[0] <= 0.2144 {
					if s[8] <= 2.2341 {
						if s[6] <= 2.8468 {
							return 0.86833
						} else {
							return 0.95556
						}
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= 0.8476 {
						return 0.00000
					} else {
						if s[8] <= -0.1419 {
							return 0.11111
						} else {
							return 0.83333
						}
					}
				}
			} else {
				if s[1] <= 3.5000 {
					if s[5] <= 1.3700 {
						if s[0] <= 0.0271 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[7] <= -0.2622 {
							return 0.90909
						} else {
							return 0.98973
						}
					}
				} else {
					if s[6] <= 1.3692 {
						if s[6] <= 1.3095 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						if s[5] <= 5.3592 {
							return 0.99941
						} else {
							return 0.99569
						}
					}
				}
			}
		}
	}
}

func decisionTreeB(s ScoreFeature) float32 {
	if s[7] <= -0.3736 {
		if s[8] <= -0.7500 {
			if s[9] <= 0.4530 {
				if s[1] <= 33.0000 {
					if s[6] <= 0.8767 {
						if s[5] <= -2.2813 {
							return 0.00026
						} else {
							return 0.06780
						}
					} else {
						if s[0] <= 0.2106 {
							return 0.67647
						} else {
							return 0.00000
						}
					}
				} else {
					if s[8] <= -5.8516 {
						if s[1] <= 70.5000 {
							return 0.47619
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[3] <= 0.5000 {
					if s[7] <= -8.4321 {
						if s[5] <= -6.8696 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				} else {
					if s[10] <= 0.5000 {
						return 1.00000
					} else {
						if s[0] <= 0.5000 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[8] <= 1.1729 {
				if s[5] <= -0.5175 {
					if s[0] <= 0.0109 {
						if s[9] <= 0.4375 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				} else {
					if s[5] <= 2.6410 {
						if s[7] <= -2.9086 {
							return 0.62500
						} else {
							return 0.93750
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= -0.4204 {
					if s[1] <= 8.0000 {
						if s[7] <= -0.9149 {
							return 0.00000
						} else {
							return 0.50000
						}
					} else {
						if s[7] <= -5.5325 {
							return 0.37500
						} else {
							return 1.00000
						}
					}
				} else {
					if s[6] <= 2.6857 {
						if s[8] <= 2.4349 {
							return 1.00000
						} else {
							return 0.84000
						}
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[5] <= -0.3786 {
			if s[7] <= 2.7475 {
				if s[7] <= 0.4749 {
					if s[8] <= 2.1609 {
						if s[0] <= 0.0579 {
							return 0.00000
						} else {
							return 0.22222
						}
					} else {
						return 1.00000
					}
				} else {
					if s[8] <= 2.2117 {
						return 0.00000
					} else {
						if s[5] <= -1.1733 {
							return 1.00000
						} else {
							return 0.00000
						}
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
			if s[6] <= 0.4591 {
				if s[6] <= -1.2175 {
					if s[1] <= 15.5000 {
						if s[1] <= 3.5000 {
							return 0.00000
						} else {
							return 0.06383
						}
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= 2.7374 {
						if s[7] <= 4.5671 {
							return 0.75362
						} else {
							return 0.00000
						}
					} else {
						if s[8] <= 1.8884 {
							return 0.16667
						} else {
							return 0.69231
						}
					}
				}
			} else {
				if s[0] <= 0.2317 {
					if s[2] <= 0.5000 {
						if s[1] <= 0.5000 {
							return 0.70213
						} else {
							return 0.96358
						}
					} else {
						if s[8] <= 2.2379 {
							return 0.97391
						} else {
							return 0.99885
						}
					}
				} else {
					if s[1] <= 2.5000 {
						if s[0] <= 0.2556 {
							return 0.13333
						} else {
							return 0.76923
						}
					} else {
						if s[8] <= -0.1419 {
							return 0.28571
						} else {
							return 0.90588
						}
					}
				}
			}
		}
	}
}

func decisionTreeC(s ScoreFeature) float32 {
	if s[7] <= 0.0820 {
		if s[0] <= 0.3011 {
			if s[9] <= 0.3961 {
				if s[8] <= -0.6857 {
					if s[8] <= -1.5803 {
						if s[6] <= 0.8876 {
							return 0.00587
						} else {
							return 0.67273
						}
					} else {
						if s[5] <= 2.6677 {
							return 0.03704
						} else {
							return 1.00000
						}
					}
				} else {
					if s[7] <= -3.7219 {
						if s[7] <= -5.4008 {
							return 0.86806
						} else {
							return 1.00000
						}
					} else {
						if s[1] <= 21.5000 {
							return 0.56383
						} else {
							return 0.96296
						}
					}
				}
			} else {
				if s[6] <= -1.4393 {
					if s[9] <= 0.8542 {
						return 1.00000
					} else {
						if s[7] <= -9.1412 {
							return 0.71429
						} else {
							return 1.00000
						}
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[9] <= 0.5625 {
				if s[5] <= -0.1999 {
					if s[6] <= -3.4710 {
						return 0.00000
					} else {
						if s[6] <= -3.4336 {
							return 1.00000
						} else {
							return 0.01429
						}
					}
				} else {
					if s[5] <= -0.1509 {
						return 1.00000
					} else {
						if s[6] <= -0.6175 {
							return 0.04587
						} else {
							return 0.42105
						}
					}
				}
			} else {
				if s[5] <= -6.9432 {
					return 0.00000
				} else {
					if s[3] <= 0.5000 {
						if s[6] <= -4.7013 {
							return 0.50000
						} else {
							return 1.00000
						}
					} else {
						return 1.00000
					}
				}
			}
		}
	} else {
		if s[1] <= 9.5000 {
			if s[8] <= 0.7421 {
				if s[2] <= 0.5000 {
					if s[6] <= -0.9220 {
						if s[1] <= 1.5000 {
							return 0.00000
						} else {
							return 0.05714
						}
					} else {
						if s[5] <= 3.6580 {
							return 0.57143
						} else {
							return 0.26316
						}
					}
				} else {
					if s[5] <= -0.2718 {
						return 0.00000
					} else {
						if s[7] <= 1.5003 {
							return 0.91667
						} else {
							return 0.56757
						}
					}
				}
			} else {
				if s[8] <= 1.7109 {
					if s[6] <= 0.0204 {
						if s[8] <= 1.3342 {
							return 1.00000
						} else {
							return 0.09091
						}
					} else {
						if s[6] <= 2.4897 {
							return 0.94000
						} else {
							return 0.79167
						}
					}
				} else {
					if s[9] <= 0.2929 {
						if s[5] <= 1.3652 {
							return 0.66667
						} else {
							return 0.97094
						}
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[5] <= -0.1999 {
				if s[8] <= 1.2148 {
					if s[6] <= -0.6579 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					return 0.00000
				}
			} else {
				if s[1] <= 16.5000 {
					if s[6] <= 1.5392 {
						if s[7] <= 0.6201 {
							return 1.00000
						} else {
							return 0.20000
						}
					} else {
						if s[1] <= 14.5000 {
							return 0.98792
						} else {
							return 0.94068
						}
					}
				} else {
					if s[6] <= 0.8061 {
						if s[1] <= 35.0000 {
							return 1.00000
						} else {
							return 0.14286
						}
					} else {
						if s[5] <= 1.7718 {
							return 0.97938
						} else {
							return 0.99882
						}
					}
				}
			}
		}
	}
}

func decisionTreeD(s ScoreFeature) float32 {
	if s[8] <= -0.1419 {
		if s[2] <= 0.5000 {
			if s[6] <= -0.4400 {
				if s[8] <= -0.5283 {
					if s[5] <= -5.5706 {
						if s[1] <= 4.5000 {
							return 0.00000
						} else {
							return 0.00478
						}
					} else {
						if s[1] <= 5.5000 {
							return 0.00526
						} else {
							return 0.18077
						}
					}
				} else {
					if s[6] <= -1.1067 {
						if s[7] <= -4.3118 {
							return 0.50000
						} else {
							return 0.00000
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[6] <= 1.7619 {
					if s[1] <= 5.5000 {
						if s[0] <= 0.0665 {
							return 0.23529
						} else {
							return 0.00000
						}
					} else {
						if s[10] <= 0.5000 {
							return 0.97059
						} else {
							return 0.29412
						}
					}
				} else {
					if s[5] <= 2.1927 {
						return 1.00000
					} else {
						if s[8] <= -8.3814 {
							return 0.48571
						} else {
							return 1.00000
						}
					}
				}
			}
		} else {
			if s[6] <= -0.2268 {
				if s[6] <= -1.3963 {
					if s[5] <= -2.7367 {
						if s[6] <= -4.2493 {
							return 0.00000
						} else {
							return 0.04138
						}
					} else {
						if s[7] <= -1.4940 {
							return 0.25490
						} else {
							return 0.00000
						}
					}
				} else {
					if s[1] <= 21.0000 {
						if s[9] <= 0.1929 {
							return 0.43243
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[1] <= 7.5000 {
					if s[1] <= 4.5000 {
						if s[5] <= -0.1745 {
							return 0.00000
						} else {
							return 0.95000
						}
					} else {
						if s[6] <= 1.1078 {
							return 0.00000
						} else {
							return 0.77778
						}
					}
				} else {
					if s[7] <= 2.8043 {
						if s[0] <= 0.0144 {
							return 0.88889
						} else {
							return 0.45455
						}
					} else {
						if s[7] <= 4.4605 {
							return 0.94231
						} else {
							return 0.99355
						}
					}
				}
			}
		}
	} else {
		if s[8] <= 0.8042 {
			if s[6] <= -0.3844 {
				if s[5] <= 1.1838 {
					return 0.00000
				} else {
					if s[7] <= 3.2234 {
						if s[8] <= 0.0908 {
							return 1.00000
						} else {
							return 0.08333
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[7] <= -0.6079 {
					if s[5] <= 2.7232 {
						if s[7] <= -8.3669 {
							return 1.00000
						} else {
							return 0.08333
						}
					} else {
						return 1.00000
					}
				} else {
					if s[4] <= 0.5000 {
						if s[5] <= 4.3390 {
							return 0.97647
						} else {
							return 0.50000
						}
					} else {
						if s[7] <= 4.5232 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[5] <= 0.9837 {
				if s[6] <= -1.2318 {
					if s[0] <= 0.1710 {
						return 1.00000
					} else {
						return 0.00000
					}
				} else {
					if s[3] <= 0.5000 {
						if s[7] <= 0.5476 {
							return 0.94444
						} else {
							return 0.33333
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[7] <= 2.1563 {
					if s[6] <= 0.3217 {
						if s[5] <= 3.0004 {
							return 0.55556
						} else {
							return 0.00000
						}
					} else {
						if s[6] <= 2.2250 {
							return 0.91818
						} else {
							return 0.99455
						}
					}
				} else {
					if s[0] <= 0.7098 {
						if s[2] <= 0.5000 {
							return 0.97628
						} else {
							return 0.99543
						}
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeE(s ScoreFeature) float32 {
	if s[7] <= -0.7861 {
		if s[5] <= -0.5565 {
			if s[6] <= 1.2051 {
				if s[5] <= -2.3519 {
					if s[5] <= -5.4142 {
						if s[5] <= -9.2655 {
							return 0.00000
						} else {
							return 0.00543
						}
					} else {
						if s[1] <= 5.5000 {
							return 0.00377
						} else {
							return 0.10658
						}
					}
				} else {
					if s[1] <= 4.5000 {
						if s[5] <= -1.4702 {
							return 0.00000
						} else {
							return 0.03846
						}
					} else {
						if s[5] <= -2.3496 {
							return 1.00000
						} else {
							return 0.28571
						}
					}
				}
			} else {
				if s[1] <= 4.0000 {
					return 0.00000
				} else {
					if s[0] <= 0.0278 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[8] <= -0.7145 {
				if s[1] <= 5.5000 {
					if s[5] <= 0.0527 {
						if s[4] <= 0.5000 {
							return 0.23810
						} else {
							return 1.00000
						}
					} else {
						if s[10] <= 0.5000 {
							return 0.25000
						} else {
							return 0.00862
						}
					}
				} else {
					if s[10] <= 0.5000 {
						return 1.00000
					} else {
						if s[1] <= 20.5000 {
							return 0.38462
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[1] <= 7.5000 {
					if s[3] <= 0.5000 {
						if s[6] <= 0.7342 {
							return 0.18182
						} else {
							return 0.83333
						}
					} else {
						if s[5] <= 2.4350 {
							return 0.41667
						} else {
							return 0.95238
						}
					}
				} else {
					if s[2] <= 1.5000 {
						if s[7] <= -2.9806 {
							return 0.95294
						} else {
							return 0.81818
						}
					} else {
						if s[8] <= 1.8436 {
							return 0.90909
						} else {
							return 1.00000
						}
					}
				}
			}
		}
	} else {
		if s[5] <= -0.3786 {
			if s[1] <= 1.5000 {
				return 0.00000
			} else {
				if s[7] <= 0.9174 {
					if s[6] <= -0.4617 {
						return 0.00000
					} else {
						if s[6] <= 2.3116 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[5] <= -1.3135 {
						if s[2] <= 1.5000 {
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
			if s[7] <= 1.4564 {
				if s[2] <= 0.5000 {
					if s[7] <= 1.3885 {
						if s[6] <= -0.4424 {
							return 0.12121
						} else {
							return 0.79412
						}
					} else {
						if s[1] <= 1.0000 {
							return 0.14286
						} else {
							return 0.50000
						}
					}
				} else {
					if s[1] <= 60.0000 {
						if s[5] <= 0.0481 {
							return 0.68750
						} else {
							return 0.91339
						}
					} else {
						if s[8] <= -2.6312 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[6] <= 0.0716 {
					if s[6] <= -2.1459 {
						return 0.00000
					} else {
						if s[8] <= -1.2564 {
							return 0.21429
						} else {
							return 0.66667
						}
					}
				} else {
					if s[8] <= 2.2566 {
						if s[4] <= 0.5000 {
							return 0.91697
						} else {
							return 0.80303
						}
					} else {
						if s[0] <= 0.6541 {
							return 0.99212
						} else {
							return 0.00000
						}
					}
				}
			}
		}
	}
}

func decisionTreeF(s ScoreFeature) float32 {
	if s[7] <= -0.3259 {
		if s[6] <= -0.4445 {
			if s[9] <= 0.4722 {
				if s[1] <= 47.5000 {
					if s[8] <= 1.3195 {
						if s[7] <= -2.0334 {
							return 0.00086
						} else {
							return 0.04206
						}
					} else {
						if s[0] <= 0.7828 {
							return 0.59091
						} else {
							return 0.00000
						}
					}
				} else {
					if s[5] <= -1.5544 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			} else {
				if s[10] <= 0.5000 {
					if s[5] <= -10.0717 {
						return 0.00000
					} else {
						if s[5] <= -1.4964 {
							return 1.00000
						} else {
							return 0.90909
						}
					}
				} else {
					if s[9] <= 0.8819 {
						return 0.00000
					} else {
						if s[6] <= -7.4853 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				}
			}
		} else {
			if s[7] <= -4.5768 {
				if s[10] <= 0.5000 {
					if s[2] <= 1.5000 {
						if s[1] <= 2.5000 {
							return 0.00000
						} else {
							return 1.00000
						}
					} else {
						if s[6] <= 4.5212 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[0] <= 0.0161 {
						if s[2] <= 0.5000 {
							return 0.03704
						} else {
							return 0.94872
						}
					} else {
						if s[5] <= -0.0862 {
							return 0.00000
						} else {
							return 0.97368
						}
					}
				}
			} else {
				if s[7] <= -2.5706 {
					if s[7] <= -3.8866 {
						if s[5] <= 2.0048 {
							return 0.25000
						} else {
							return 1.00000
						}
					} else {
						if s[8] <= 6.6486 {
							return 0.05263
						} else {
							return 1.00000
						}
					}
				} else {
					if s[5] <= 0.0130 {
						if s[7] <= -0.8949 {
							return 0.13333
						} else {
							return 1.00000
						}
					} else {
						if s[8] <= -1.0320 {
							return 0.50000
						} else {
							return 0.95652
						}
					}
				}
			}
		}
	} else {
		if s[6] <= -0.1516 {
			if s[6] <= -1.2978 {
				if s[1] <= 16.5000 {
					if s[7] <= 3.2710 {
						if s[2] <= 0.5000 {
							return 0.00000
						} else {
							return 0.04000
						}
					} else {
						if s[8] <= 1.4036 {
							return 0.05882
						} else {
							return 0.60000
						}
					}
				} else {
					if s[3] <= 0.5000 {
						if s[1] <= 23.5000 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[0] <= 0.2059 {
					if s[1] <= 6.5000 {
						if s[6] <= -0.8843 {
							return 0.35714
						} else {
							return 0.73913
						}
					} else {
						return 1.00000
					}
				} else {
					if s[6] <= -0.9090 {
						return 1.00000
					} else {
						if s[0] <= 0.4428 {
							return 0.00000
						} else {
							return 0.44444
						}
					}
				}
			}
		} else {
			if s[0] <= 0.2317 {
				if s[9] <= 0.1010 {
					if s[6] <= 0.6980 {
						if s[8] <= -0.3788 {
							return 0.35294
						} else {
							return 0.84848
						}
					} else {
						if s[1] <= 0.5000 {
							return 0.87654
						} else {
							return 0.98597
						}
					}
				} else {
					if s[5] <= 0.8536 {
						return 1.00000
					} else {
						if s[9] <= 0.2429 {
							return 0.16667
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[8] <= 0.1861 {
					if s[6] <= 2.1761 {
						if s[1] <= 0.5000 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						return 1.00000
					}
				} else {
					if s[9] <= 0.1500 {
						if s[3] <= 0.5000 {
							return 0.75410
						} else {
							return 0.97297
						}
					} else {
						return 0.00000
					}
				}
			}
		}
	}
}

func decisionTreeG(s ScoreFeature) float32 {
	if s[6] <= -0.4424 {
		if s[6] <= -1.5772 {
			if s[6] <= -3.8231 {
				if s[5] <= -5.1702 {
					if s[4] <= 0.5000 {
						if s[3] <= 0.5000 {
							return 0.00000
						} else {
							return 0.00062
						}
					} else {
						return 0.00000
					}
				} else {
					if s[9] <= 0.5179 {
						if s[2] <= 1.5000 {
							return 0.00332
						} else {
							return 0.04545
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[8] <= 1.6541 {
					if s[10] <= 0.5000 {
						if s[0] <= 0.7298 {
							return 0.18898
						} else {
							return 0.66667
						}
					} else {
						if s[8] <= -0.5254 {
							return 0.01552
						} else {
							return 0.08108
						}
					}
				} else {
					if s[8] <= 2.3333 {
						return 1.00000
					} else {
						if s[3] <= 0.5000 {
							return 0.66667
						} else {
							return 0.50000
						}
					}
				}
			}
		} else {
			if s[9] <= 0.2792 {
				if s[0] <= 0.6190 {
					if s[8] <= -0.7482 {
						if s[7] <= -0.0457 {
							return 0.04420
						} else {
							return 0.25926
						}
					} else {
						if s[5] <= 4.0456 {
							return 0.31579
						} else {
							return 1.00000
						}
					}
				} else {
					if s[7] <= -2.1992 {
						return 0.00000
					} else {
						return 1.00000
					}
				}
			} else {
				if s[10] <= 0.5000 {
					return 1.00000
				} else {
					if s[6] <= -1.2761 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[5] <= -0.1999 {
			if s[10] <= 0.5000 {
				if s[1] <= 2.0000 {
					return 0.00000
				} else {
					return 1.00000
				}
			} else {
				if s[2] <= 0.5000 {
					if s[7] <= -0.8078 {
						return 0.00000
					} else {
						if s[7] <= 0.3228 {
							return 1.00000
						} else {
							return 0.25000
						}
					}
				} else {
					if s[6] <= 1.7146 {
						if s[1] <= 25.0000 {
							return 0.80000
						} else {
							return 0.00000
						}
					} else {
						return 1.00000
					}
				}
			}
		} else {
			if s[1] <= 10.5000 {
				if s[8] <= 0.6345 {
					if s[0] <= 0.2226 {
						if s[6] <= 2.9036 {
							return 0.65493
						} else {
							return 0.91228
						}
					} else {
						if s[8] <= -1.0256 {
							return 0.00000
						} else {
							return 0.50000
						}
					}
				} else {
					if s[8] <= 2.1594 {
						if s[6] <= 0.0315 {
							return 0.50000
						} else {
							return 0.92357
						}
					} else {
						if s[7] <= 0.1039 {
							return 0.73810
						} else {
							return 0.97942
						}
					}
				}
			} else {
				if s[6] <= 0.7694 {
					if s[3] <= 0.5000 {
						if s[6] <= 0.6342 {
							return 0.52941
						} else {
							return 0.00000
						}
					} else {
						return 1.00000
					}
				} else {
					if s[5] <= 8.5196 {
						if s[3] <= 0.5000 {
							return 0.99563
						} else {
							return 0.98924
						}
					} else {
						if s[7] <= 8.7842 {
							return 0.75000
						} else {
							return 1.00000
						}
					}
				}
			}
		}
	}
}

func decisionTreeH(s ScoreFeature) float32 {
	if s[7] <= -0.7551 {
		if s[6] <= -0.4208 {
			if s[9] <= 0.4643 {
				if s[0] <= 0.8590 {
					if s[8] <= 1.5740 {
						if s[9] <= 0.2792 {
							return 0.00375
						} else {
							return 0.28571
						}
					} else {
						if s[1] <= 4.0000 {
							return 0.71429
						} else {
							return 0.25000
						}
					}
				} else {
					if s[8] <= -18.4861 {
						return 0.00000
					} else {
						if s[5] <= -3.6223 {
							return 0.00000
						} else {
							return 0.02105
						}
					}
				}
			} else {
				if s[6] <= -7.1628 {
					if s[9] <= 0.9167 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.5455 {
						if s[1] <= 8.5000 {
							return 1.00000
						} else {
							return 0.95652
						}
					} else {
						if s[8] <= -5.3725 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			}
		} else {
			if s[1] <= 5.5000 {
				if s[10] <= 0.5000 {
					if s[2] <= 0.5000 {
						return 1.00000
					} else {
						if s[7] <= -6.8249 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				} else {
					if s[3] <= 0.5000 {
						if s[2] <= 0.5000 {
							return 0.05634
						} else {
							return 0.66667
						}
					} else {
						if s[6] <= 1.7622 {
							return 0.27778
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[9] <= 0.1709 {
					if s[0] <= 0.2208 {
						if s[3] <= 0.5000 {
							return 0.84810
						} else {
							return 0.96970
						}
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
			}
		}
	} else {
		if s[8] <= 0.8620 {
			if s[6] <= -0.9239 {
				if s[1] <= 2.5000 {
					if s[6] <= -0.9908 {
						return 0.00000
					} else {
						if s[5] <= 0.9203 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[5] <= 2.5384 {
						if s[8] <= -3.0042 {
							return 0.03226
						} else {
							return 0.50000
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[2] <= 0.5000 {
					if s[1] <= 4.5000 {
						if s[0] <= 0.1989 {
							return 0.70588
						} else {
							return 0.16667
						}
					} else {
						if s[7] <= 3.0989 {
							return 0.88571
						} else {
							return 0.46154
						}
					}
				} else {
					if s[5] <= -0.5670 {
						if s[5] <= -1.3135 {
							return 0.75000
						} else {
							return 0.00000
						}
					} else {
						if s[6] <= 0.5801 {
							return 0.52632
						} else {
							return 0.97030
						}
					}
				}
			}
		} else {
			if s[0] <= 0.5603 {
				if s[0] <= 0.2208 {
					if s[1] <= 8.5000 {
						if s[0] <= 0.0232 {
							return 0.90759
						} else {
							return 0.99038
						}
					} else {
						if s[1] <= 19.5000 {
							return 0.98797
						} else {
							return 0.99784
						}
					}
				} else {
					if s[1] <= 5.5000 {
						if s[7] <= 1.6437 {
							return 0.50000
						} else {
							return 0.88889
						}
					} else {
						if s[0] <= 0.2232 {
							return 0.00000
						} else {
							return 0.98246
						}
					}
				}
			} else {
				return 0.00000
			}
		}
	}
}

func decisionTreeI(s ScoreFeature) float32 {
	if s[6] <= -0.4300 {
		if s[7] <= -2.0334 {
			if s[9] <= 0.5357 {
				if s[1] <= 7.5000 {
					if s[9] <= 0.4226 {
						if s[3] <= 0.5000 {
							return 0.00010
						} else {
							return 0.00148
						}
					} else {
						if s[5] <= -6.0521 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[5] <= -1.6698 {
						if s[6] <= -3.4657 {
							return 0.00050
						} else {
							return 0.07143
						}
					} else {
						if s[9] <= 0.0400 {
							return 0.40000
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[6] <= -7.6342 {
					if s[9] <= 0.9583 {
						return 0.00000
					} else {
						if s[1] <= 6.5000 {
							return 0.00000
						} else {
							return 1.00000
						}
					}
				} else {
					if s[9] <= 0.8661 {
						return 1.00000
					} else {
						if s[9] <= 0.8819 {
							return 0.40000
						} else {
							return 1.00000
						}
					}
				}
			}
		} else {
			if s[6] <= -2.0305 {
				if s[8] <= 6.8227 {
					if s[6] <= -2.5739 {
						if s[5] <= -4.6595 {
							return 0.05000
						} else {
							return 0.00000
						}
					} else {
						if s[7] <= -1.4138 {
							return 0.18750
						} else {
							return 0.02381
						}
					}
				} else {
					return 1.00000
				}
			} else {
				if s[7] <= 3.2234 {
					if s[1] <= 6.5000 {
						if s[8] <= -2.3177 {
							return 0.02439
						} else {
							return 0.37143
						}
					} else {
						if s[1] <= 13.5000 {
							return 0.62963
						} else {
							return 0.00000
						}
					}
				} else {
					if s[8] <= 3.6313 {
						if s[0] <= 0.0833 {
							return 0.80000
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[1] <= 6.5000 {
			if s[5] <= 1.3652 {
				if s[8] <= -0.6870 {
					if s[5] <= -2.8317 {
						if s[5] <= -4.2410 {
							return 1.00000
						} else {
							return 0.80000
						}
					} else {
						if s[1] <= 5.5000 {
							return 0.10870
						} else {
							return 0.87500
						}
					}
				} else {
					if s[6] <= 2.2950 {
						if s[6] <= 1.6996 {
							return 0.67742
						} else {
							return 1.00000
						}
					} else {
						return 0.00000
					}
				}
			} else {
				if s[7] <= 0.4564 {
					if s[8] <= -8.6008 {
						if s[2] <= 0.5000 {
							return 0.08333
						} else {
							return 1.00000
						}
					} else {
						if s[7] <= -4.8800 {
							return 0.96552
						} else {
							return 0.51064
						}
					}
				} else {
					if s[0] <= 0.4490 {
						if s[0] <= 0.2313 {
							return 0.95794
						} else {
							return 0.79310
						}
					} else {
						return 0.00000
					}
				}
			}
		} else {
			if s[5] <= -0.4322 {
				if s[10] <= 0.5000 {
					return 1.00000
				} else {
					if s[7] <= -1.3413 {
						return 0.00000
					} else {
						if s[5] <= -1.2281 {
							return 1.00000
						} else {
							return 0.00000
						}
					}
				}
			} else {
				if s[9] <= 0.0800 {
					if s[8] <= -0.3734 {
						if s[2] <= 0.5000 {
							return 0.66667
						} else {
							return 0.93885
						}
					} else {
						if s[0] <= 0.1988 {
							return 0.99641
						} else {
							return 0.95294
						}
					}
				} else {
					if s[9] <= 0.3079 {
						if s[7] <= 2.2670 {
							return 0.88889
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
}

func decisionTreeJ(s ScoreFeature) float32 {
	if s[5] <= -0.4300 {
		if s[6] <= -0.4097 {
			if s[6] <= -3.8234 {
				if s[9] <= 0.7750 {
					if s[7] <= -0.2599 {
						if s[0] <= 0.0064 {
							return 0.00073
						} else {
							return 0.00015
						}
					} else {
						if s[0] <= 0.0464 {
							return 0.33333
						} else {
							return 0.00000
						}
					}
				} else {
					if s[3] <= 0.5000 {
						if s[10] <= 0.5000 {
							return 1.00000
						} else {
							return 0.00000
						}
					} else {
						return 1.00000
					}
				}
			} else {
				if s[9] <= 0.3961 {
					if s[0] <= 0.9510 {
						if s[2] <= 0.5000 {
							return 0.00131
						} else {
							return 0.05181
						}
					} else {
						if s[5] <= -3.1460 {
							return 0.66667
						} else {
							return 0.00000
						}
					}
				} else {
					return 1.00000
				}
			}
		} else {
			if s[10] <= 0.5000 {
				return 1.00000
			} else {
				if s[2] <= 0.5000 {
					if s[6] <= 0.9996 {
						return 0.00000
					} else {
						if s[7] <= -0.9989 {
							return 0.00000
						} else {
							return 0.60000
						}
					}
				} else {
					if s[0] <= 0.0109 {
						return 1.00000
					} else {
						return 0.00000
					}
				}
			}
		}
	} else {
		if s[0] <= 0.2317 {
			if s[5] <= 1.6833 {
				if s[6] <= -1.3243 {
					return 0.00000
				} else {
					if s[2] <= 0.5000 {
						if s[0] <= 0.0313 {
							return 0.65909
						} else {
							return 0.90000
						}
					} else {
						if s[7] <= 1.0572 {
							return 0.89394
						} else {
							return 1.00000
						}
					}
				}
			} else {
				if s[7] <= 0.4238 {
					if s[6] <= 2.3802 {
						if s[1] <= 6.5000 {
							return 0.27586
						} else {
							return 0.68966
						}
					} else {
						if s[8] <= 4.1779 {
							return 0.96396
						} else {
							return 1.00000
						}
					}
				} else {
					if s[5] <= 6.1870 {
						if s[6] <= 0.0541 {
							return 0.35714
						} else {
							return 0.98074
						}
					} else {
						if s[8] <= -4.0916 {
							return 0.96154
						} else {
							return 1.00000
						}
					}
				}
			}
		} else {
			if s[6] <= 0.5766 {
				if s[0] <= 0.8655 {
					if s[8] <= -1.1937 {
						if s[5] <= -0.1509 {
							return 0.62500
						} else {
							return 0.05556
						}
					} else {
						if s[5] <= 1.9653 {
							return 0.00000
						} else {
							return 0.51786
						}
					}
				} else {
					return 0.00000
				}
			} else {
				if s[5] <= 1.7158 {
					if s[0] <= 0.5451 {
						return 0.00000
					} else {
						return 1.00000
					}
				} else {
					if s[0] <= 0.2352 {
						if s[3] <= 0.5000 {
							return 0.30000
						} else {
							return 1.00000
						}
					} else {
						if s[6] <= 1.9192 {
							return 0.78571
						} else {
							return 0.97436
						}
					}
				}
			}
		}
	}
}
