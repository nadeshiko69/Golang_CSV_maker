FROM golang:1.25.6-bookworm

# 非rootユーザーを作成
RUN groupadd --gid 1000 vscode \
    && useradd --uid 1000 --gid 1000 -m -s /bin/bash vscode \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
       sudo \
       bash-completion \
       ca-certificates \
       git \
       less \
    && echo "vscode ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/vscode \
    && chmod 0440 /etc/sudoers.d/vscode \
    && rm -rf /var/lib/apt/lists/*

SHELL ["/bin/bash", "-lc"]

RUN grep -q "bash-completion/bash_completion" /etc/bash.bashrc || \
    echo '[ -f /usr/share/bash-completion/bash_completion ] && . /usr/share/bash-completion/bash_completion' >> /etc/bash.bashrc

WORKDIR /usr/src/app

COPY go.mod ./
# RUN go mod download

COPY . .

RUN chown -R vscode:vscode /usr/src/app

# A Tour Of Go環境の導入
# CMD ["go", "install", "golang.org/x/website/tour@latest"]
# ENV PATH="/home/vscode/go/bin:${PATH}"

# 普段のシェルを vscode ユーザーの bash に
USER vscode
ENV SHELL=/bin/bash