package btree

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestDelete1(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("1"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete1() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete1() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=1 cap=1) {
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=2 cap=2) {
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "12",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "15",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=8) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "10",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "11",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "13",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "14",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=4) {
       (btree.keyStruct) {
        key: (string) (len=2) "16",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "17",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "18",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   }),
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "3",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=4) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "2",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "20",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "4",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "5",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "6",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "8",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "9",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete2(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("2"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete2() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete2() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=1 cap=1) {
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=2 cap=2) {
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "12",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "15",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=8) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "1",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "10",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "11",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "13",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "14",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=4) {
       (btree.keyStruct) {
        key: (string) (len=2) "16",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "17",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "18",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   }),
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "4",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=4) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=2) "20",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "3",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "5",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "6",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "8",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "9",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete3(t *testing.T) {
	// should require combining of nodes
	B := setupDeleteTest()

	if _, err := B.Delete("3"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete3() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete3() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=1 cap=1) {
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=2 cap=2) {
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "12",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "15",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=8) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "1",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "10",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "11",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "13",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "14",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=4) {
       (btree.keyStruct) {
        key: (string) (len=2) "16",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "17",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "18",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   }),
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "4",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=4) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "2",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "20",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "5",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "6",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "8",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "9",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete5(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("5"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete5() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete5() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=1 cap=1) {
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=2 cap=2) {
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "12",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "15",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=8) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "1",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "10",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "11",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "13",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "14",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=4) {
       (btree.keyStruct) {
        key: (string) (len=2) "16",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "17",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "18",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   }),
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "3",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=4) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "2",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "20",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "4",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "6",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "8",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "9",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete6(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("6"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete6() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete6() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=1 cap=1) {
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=2 cap=2) {
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "12",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "15",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=8) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "1",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "10",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "11",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "13",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "14",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=4) {
       (btree.keyStruct) {
        key: (string) (len=2) "16",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "17",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "18",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   }),
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "3",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=4) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "2",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "20",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "4",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "5",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "8",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "9",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete12(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("12"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete12() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete12() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=1 cap=1) {
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=2 cap=2) {
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "11",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "15",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=8) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "1",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "10",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "13",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "14",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=4) {
       (btree.keyStruct) {
        key: (string) (len=2) "16",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "17",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "18",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   }),
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "3",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=4) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "2",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "20",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "4",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "5",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "6",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "8",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "9",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete12then11(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("12"); err != nil {
		t.Fatal(err)
	}
	if _, err := B.Delete("11"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete12then11() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete12then11() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=1 cap=1) {
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=2 cap=2) {
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "13",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "16",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=8) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "1",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "10",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=2) "14",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "15",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=4) {
       (btree.keyStruct) {
        key: (string) (len=2) "17",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "18",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   }),
   (*btree.node)({
    leaf: (bool) false,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "3",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) (len=3 cap=4) {
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "2",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=2) "20",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=3 cap=8) {
       (btree.keyStruct) {
        key: (string) (len=1) "4",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "5",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "6",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     }),
     (*btree.node)({
      leaf: (bool) true,
      keys: ([]btree.keyStruct) (len=2 cap=2) {
       (btree.keyStruct) {
        key: (string) (len=1) "8",
        rowPtr: (int64) 0
       },
       (btree.keyStruct) {
        key: (string) (len=1) "9",
        rowPtr: (int64) 0
       }
      },
      child: ([]*btree.node) {
      }
     })
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete12then11then15(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("12"); err != nil {
		t.Fatal(err)
	}
	if _, err := B.Delete("11"); err != nil {
		t.Fatal(err)
	}
	if _, err := B.Delete("15"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete12then11then15() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete12then11then15() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=4 cap=8) {
   (btree.keyStruct) {
    key: (string) (len=2) "16",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=1) "3",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=1) "7",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=5 cap=8) {
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=4 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=1) "1",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "10",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "13",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "14",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=4) {
     (btree.keyStruct) {
      key: (string) (len=2) "17",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "18",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "2",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "20",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=3 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=1) "4",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "5",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "6",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "8",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "9",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete12then11then1(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("12"); err != nil {
		t.Fatal(err)
	}
	if _, err := B.Delete("11"); err != nil {
		t.Fatal(err)
	}
	if _, err := B.Delete("1"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete12then11then1() {
		t.Fatal("Tree result did not match")
	}
}

func ResultDelete12then11then1() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=4 cap=8) {
   (btree.keyStruct) {
    key: (string) (len=2) "16",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=1) "3",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=1) "7",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=5 cap=8) {
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=4 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "10",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "13",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "14",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "15",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=4) {
     (btree.keyStruct) {
      key: (string) (len=2) "17",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "18",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "2",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "20",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=3 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=1) "4",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "5",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "6",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "8",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "9",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDelete4then8(t *testing.T) {
	B := setupDeleteTest()

	if _, err := B.Delete("4"); err != nil {
		t.Fatal(err)
	}

	if _, err := B.Delete("8"); err != nil {
		t.Fatal(err)
	}

	if spew.Sdump(B) != ResultDelete4then8() {
		t.Fatal("Tree result did not match")
	}

}

func ResultDelete4then8() string {
	return `(*btree.Btree)({
 root: (*btree.node)({
  leaf: (bool) false,
  keys: ([]btree.keyStruct) (len=4 cap=8) {
   (btree.keyStruct) {
    key: (string) (len=2) "12",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=2) "15",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=2) "19",
    rowPtr: (int64) 0
   },
   (btree.keyStruct) {
    key: (string) (len=1) "3",
    rowPtr: (int64) 0
   }
  },
  child: ([]*btree.node) (len=5 cap=8) {
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=3 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=1) "1",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "10",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "11",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=2) "13",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "14",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=3 cap=4) {
     (btree.keyStruct) {
      key: (string) (len=2) "16",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "17",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "18",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=2 cap=2) {
     (btree.keyStruct) {
      key: (string) (len=1) "2",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=2) "20",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   }),
   (*btree.node)({
    leaf: (bool) true,
    keys: ([]btree.keyStruct) (len=4 cap=8) {
     (btree.keyStruct) {
      key: (string) (len=1) "5",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "6",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "7",
      rowPtr: (int64) 0
     },
     (btree.keyStruct) {
      key: (string) (len=1) "9",
      rowPtr: (int64) 0
     }
    },
    child: ([]*btree.node) {
    }
   })
  }
 }),
 maxElements: (int) 4,
 minElements: (int) 2
})
`
}

func TestDeleteLastElement(t *testing.T) {
	B := New(5)

	B.Insert("1", 0)

	_, err := B.Delete("1")

	if err != nil {
		t.Fatal(err)
	}
}

func setupDeleteTest() *Btree {
	B := New(5)

	data := []string{"14", "7", "12", "3", "4", "10", "19", "11", "1", "2", "20", "6", "17", "13", "9", "8", "18", "5", "15", "16"}

	for _, item := range data {
		B.Insert(item, 0)
	}

	spew.Config.DisablePointerAddresses = true

	return B
}
