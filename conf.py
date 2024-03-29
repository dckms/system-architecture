# Configuration file for the Sphinx documentation builder.
#
# This file only contains a selection of the most common options. For a full
# list see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Path setup --------------------------------------------------------------

# If extensions (or modules to document with autodoc) are in another directory,
# add these directories to sys.path here. If the directory is relative to the
# documentation root, use os.path.abspath to make it absolute, like shown here.
#
import os
# import sys
import datetime
import posixpath
from markupsafe import escape

# sys.path.insert(0, os.path.abspath('.'))
PROJECT_ROOT = os.path.dirname(os.path.abspath(__file__))
BASE_URL = os.environ.get('BASE_URL', 'http://YOUR_HOST_URL')

# -- Project information -----------------------------------------------------

project = 'System Architecture'
author = '@dckms'
copyright = '{}, {} (<a href="{}" rel="nofollow">License</a>)'.format(
    datetime.datetime.now().strftime('%Y'),
    author,
    escape(posixpath.join(BASE_URL, 'LICENSE.txt'))
)

# -- General configuration ---------------------------------------------------

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    'sphinx.ext.intersphinx',
    'sphinx.ext.todo',
    'myst_parser',
    'yasfb',
    'sphinx_reredirects',
]

source_suffix = {
    '.rst': 'restructuredtext',
    # '.txt': 'markdown',
    '.md': 'markdown',
}

# Add any paths that contain templates here, relative to this directory.
templates_path = ['_templates']

# The language for content autogenerated by Sphinx. Refer to documentation
# for a list of supported languages.
#
# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
language = 'ru'

# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
# This pattern also affects html_static_path and html_extra_path.
exclude_patterns = ['.*', '_build', 'Thumbs.db', '.DS_Store']

smartquotes_action = 'd'

# -- Options for HTML output -------------------------------------------------

html_title = project

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme = 'alabaster'

# Add any paths that contain custom static files (such as style sheets) here,
# relative to this directory. They are copied after the builtin static files,
# so a file named "default.css" will overwrite the builtin "default.css".
html_static_path = ['_static']

html_extra_path = [
    '_html_extra',
    'LICENSE.txt',
    'LICENSE-CONTENT.txt',
    'LICENSE-CODE.txt',
    'AUTHORS.txt',
]

# Do not use html_baseurl together with custom page metadata ``canonical-url``
# https://www.sphinx-doc.org/en/master/development/theming.html#use-custom-page-metadata-in-html-templates
# html_baseurl = BASE_URL

html_theme_options = {
    'github_banner': True,
    'github_button': True,
    'github_type': "star",
    'github_count': "true",
    'github_user': "dckms",
    'github_repo': "system-architecture",
    'show_related': True,
    'show_relbars': True,
    'show_powered_by': True,
    'description': "Distributed Collaborative Knowledge Management System for System Architecture",
    # Google Analytics ID
    'analytics_id': os.environ.get('GOOGLE_ANALITICS_ID', None),
    'extra_nav_links': {
        '<i class="fab fa-telegram" style="color:#54a9eb;"></i> Telegram': 'https://t.me/emacsway_log',
    },

}

html_context = {
    # Yandex.Metrika ID
    'theme_yandex_metrika_id': os.environ.get('YANDEX_METRIKA_ID', None),
}


# This is also used if you do content translation via gettext catalogs.
# Usually you set "language" from the command line for these cases.
# language = 'ru'

# This defaults to the global language selected with language.
# https://www.sphinx-doc.org/en/master/usage/configuration.html#confval-html_search_language
# html_search_language = 'ru'

# https://www.sphinx-doc.org/ru/master/usage/advanced/intl.html
locale_dirs = [os.path.join(PROJECT_ROOT, 'locale/')]
gettext_compact = False

show_authors = True

# RSS settings (yasfb)
feed_base_url = BASE_URL
feed_author = author
