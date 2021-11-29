{ overlays ? [ ]
}@args:

let
  inherit (nixpkgs) pkgs;

  env-overlay = self: super: {
    my-env = super.buildEnv {
      name = "my-env";
      paths = with self; [
        go

        entr
        gnumake
      ];
    };
  };

  nixpkgs = import <nixpkgs> (args // {
    overlays = overlays ++ [
      env-overlay
    ];
  });

in

pkgs.mkShell {
  buildInputs = with pkgs; [
    go
  ];
}
