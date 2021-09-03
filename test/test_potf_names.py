from model import get_potf_last_name


def test_real():
    assert get_potf_last_name("Markus") == "Kaarlonen"
    assert get_potf_last_name("Marko") == "Saaresto"


def test_wrong():
    assert get_potf_last_name("Jane") == "Doe"

