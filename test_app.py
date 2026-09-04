from app import divide, double


def test_double_positive():
    assert double(2) == 4


def test_double_zero():
    assert double(0) == 0


def test_double_negative():
    assert double(-3) == -6


def test_divide_positive():
    assert divide(6, 3) == 2


def test_divide_negative():
    assert divide(-6, 3) == -2


def test_divide_float_result():
    assert divide(1, 4) == 0.25
