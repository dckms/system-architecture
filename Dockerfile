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

# Set the working directory to our repo root
WORKDIR /sphinxtechnicalwriting

# Only copy the Pipfile
# https://github.com/ArtFlag/sphinxtechnicalwriting/blob/master/Pipfile
# create empty Pipfile.lock too
COPY Pipfile Pipfile.lock /sphinxtechnicalwriting/

ADD requirements.freeze.txt /sphinxtechnicalwriting
RUN pip3 install -r requirements.freeze.txt

# Install the packages
RUN pipenv install --system --deploy