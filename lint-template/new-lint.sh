#!/bin/bash
# set -x

echo Type: [`ls linters`]
read t
existing=`ls linters/$t`
echo "Exising linters:"
echo $existing
echo "Short name:"
read name

cp -r lint-template/FIXME linters/$t/

mv linters/$t/FIXME linters/$t/$name

mv linters/$t/$name/FIXME.go linters/$t/$name/$name.go
sed -i "s/package FIXME/package $name/g" linters/$t/$name/$name.go
sed -i "s/Name = \"FIXME\"/Name = \"${name}\"/g" linters/$t/$name/$name.go

mv linters/$t/$name/FIXME_test.go linters/$t/$name/${name}_test.go
sed -i "s/package FIXME_test/package ${name}_test/g" linters/$t/$name/${name}_test.go
sed -i "s|\"github.com/gibizer/operator-lint/lint-template/FIXME\"|\"github.com/gibizer/operator-lint/linters/$t/$name\"|g" linters/$t/$name/${name}_test.go
sed -i "s/TestFIXME/Test${name}/g" linters/$t/$name/${name}_test.go
sed -i "s/FIXME.NewAnalyzer()/${name}.NewAnalyzer()/g" linters/$t/$name/${name}_test.go
