{ pkgs ? import <nixpkgs> {
    config = {
      allowUnfree = true;
    };
  } }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # mysql84
  ];

  name = "go-bookstore-env";

  shellHook = ''
    export PS1="\[\033[01;32m\][$name:\\w]\$\[\033[00m\] "
    clear
    echo "on an environment."
  '';
}
