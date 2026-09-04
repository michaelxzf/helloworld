from app import divide, double, read_file


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


def test_read_file_returns_contents(tmp_path):
    file_path = tmp_path / "example.txt"
    file_path.write_text("hello world")
    assert read_file(str(file_path)) == "hello world"


def test_read_file_empty_file(tmp_path):
    file_path = tmp_path / "empty.txt"
    file_path.write_text("")
    assert read_file(str(file_path)) == ""


def test_read_file_preserves_multiline_contents(tmp_path):
    file_path = tmp_path / "lines.txt"
    file_path.write_text("first\nsecond\n")
    assert read_file(str(file_path)) == "first\nsecond\n"
