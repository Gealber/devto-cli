# file: devto
# devto parameter completation


_complete_devto () {
    local cmd="${1##*/}"
    local word=${COMP_WORDS[COMP_CWORD]}
    local xpat
    local commands="articles auth comments followers listings organizations podcasts reading_lists tags webhooks"

    #articles
    local art_subcommands="create update latest videos me"
    #listings
    local list_subcommands="create"
    #tags
    local tags_subcommands="follows"
    #webhooks
    local webhooks_subcommands="create delete"

    # Check to see what command is been executed
    cur_command=$(_find_command)
    case "$cmd" in
    devto)
        if [[ "$(_get_argc)" -eq 1 ]]; then
            xpat="$commands"
        elif [[ "$(_get_argc)" -ge 2 ]]; then
            case "$cur_command" in
            # articles subcommands
            articles)
                xpat="$art_subcommands"
                ;;

            # listings subcommands
            listings)
                xpat="$list_subcommands"
                ;;

            # tags subcommands
            tags)
                xpat="$tags_subcommands"
                ;;

            # webhooks subcommands
            webhooks)
                xpat="$webhooks_subcommands"
                ;;
            *)
                if [[ "$(_get_argc)" -eq 2 ]]; then
                    xpat="$commands"
                fi
                ;;
            esac
        fi
        ;;

    esac

    COMPREPLY=($(compgen -W "$xpat" -- "${word}"))
}

_find_command () {
    local line=${COMP_LINE}
    IFS=' ' read -r -a array <<< "$line"
    len="${#array[@]}"
    
    # check if there's no subcommand
    if [[ $len -ge 2 ]]; then
        echo "${array[1]}"
    fi
}

_get_argc () {
    local line=${COMP_LINE}
    IFS=' ' read -r -a array <<< "$line"
    lenArr="${#array[@]}"
    echo "$lenArr"
}

complete -F _complete_devto devto
