
create:
	helm create mychart

clean-template:
	rm -rf mychart/templates/*


mychart: create
	helm install clunky-serval ./mychart

show:
	helm get manifest full-coral

show2:
	helm get manifest clunky-serval

set-value:
	helm install solid-vulture ./mychart --dry-run --debug --set favoriteDrink=slurm