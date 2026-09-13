import os
import shutil
import uuid

import pytest

from app import divide, double, read_file

WORKSPACE_ROOT = os.path.dirname(os.path.abspath(__file__))


@pytest.fixture
def workspace_dir():
    d = os.path.join(WORKSPACE_ROOT, f".tmp_test_{uuid.uuid4().hex[:8]}")
    os.makedirs(d)
    yield d
    shutil.rmtree(d)


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


def test_divide_by_zero_raises_value_error():
    with pytest.raises(ValueError):
        divide(6, 0)


def test_read_file_returns_contents(workspace_dir):
    file_path = os.path.join(workspace_dir, "example.txt")
    with open(file_path, "w") as f:
        f.write("hello world")
    assert read_file(os.path.join(workspace_dir, "example.txt")) == "hello world"


def test_read_file_empty_file(workspace_dir):
    file_path = os.path.join(workspace_dir, "empty.txt")
    with open(file_path, "w") as f:
        f.write("")
    assert read_file(file_path) == ""


def test_read_file_preserves_multiline_contents(workspace_dir):
    file_path = os.path.join(workspace_dir, "lines.txt")
    with open(file_path, "w") as f:
        f.write("first\nsecond\n")
    assert read_file(file_path) == "first\nsecond\n"


def test_read_file_accepts_relative_path_in_workspace():
    assert read_file("app.py").startswith("import os")


def test_read_file_rejects_path_traversal():
    with pytest.raises(ValueError):
        read_file("../secret.txt")


def test_read_file_rejects_absolute_path_outside_workspace():
    with pytest.raises(ValueError):
        read_file("/etc/passwd")


def test_read_file_rejects_symlink_escape(workspace_dir):
    link_path = os.path.join(workspace_dir, "escape_link")
    os.symlink("/etc/passwd", link_path)
    with pytest.raises(ValueError):
        read_file(link_path)
