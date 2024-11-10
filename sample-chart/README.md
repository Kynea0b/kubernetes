## helmの使い方

helmの使い方

`help install`でtemplateにあるyamlファイルがkuberentes環境に取り込まれる・
installされたファイルを表示したい時は、`helm get manifest full-coral`する。

https://helm.sh/docs/chart_template_guide/getting_started/

### values.yaml

valueに設定した値を`configmap.yaml`で使用できる。install時に使用することもできる。

https://helm.sh/docs/chart_template_guide/values_files/