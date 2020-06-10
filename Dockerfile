FROM scratch

COPY ./src/loop ./

ENTRYPOINT ["./loop"]