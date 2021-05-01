if [[ -z $DEVTO_DIR ]]; then
	mkdir -p $HOME/.devto && cp bash/* $HOME/.devto
    DEVTO_DIR=$HOME/.devto
    cat >> $HOME/.bashrc << EOF

# >>> devto initialization >>>
# Configuration of devto for autocompletion
export DEVTO_DIR=\$HOME/.devto

if [ -f \$DEVTO_DIR/devto.sh ]; then
    . \$DEVTO_DIR/devto.sh
fi

# <<< devto initialization <<<

EOF
fi
