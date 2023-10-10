FROM python:3.8-slim-buster

# https://sphinxtechnicalwriting.readthedocs.io/en/latest/cd/config-docker.html

RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get install -y --no-install-recommends make && \
    apt-get clean && \
    apt-get install -y --no-install-recommends git && \
    rm -rf /var/lib/apt/lists/* && \
    pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir pipenv

WORKDIR /sphinxtechnicalwriting

# https://github.com/ArtFlag/sphinxtechnicalwriting/blob/master/Pipfile
RUN echo '[[source]]\nname = "pypi"\nurl = "https://pypi.org/simple"\nverify_ssl = true\n' \
    '[dev-packages]\nrstcheck = "*"\nblack = "==19.10b0"\nflake8 = "*"\ndoc8 = "*"\n' \
    '[packages]\nsphinxcontrib-mermaid = "*"\nSphinx = "3.2.1"\nsphinx-rtd-theme = "*"\n' \
    '[requires]\npython_version = "3.8.2"' > Pipfile

RUN touch Pipfile.lock

ADD requirements.freeze.txt /sphinxtechnicalwriting
RUN pip3 install -r requirements.freeze.txt

RUN pipenv install --system --deploy
