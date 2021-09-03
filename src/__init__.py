_POTF_FIRST_LAST_NAME = {
    "Markus": "Kaarlonen",
    "Marko": "Saaresto",
    "Olli": "Tukiainen",
    "Jani": "Snellman",
    "Jaska": "Mäkinen",
    "Jari": "Salminen",
}

def get_potf_last_name(first_name):
    return _POTF_FIRST_LAST_NAME.get(first_name, "Doe")

