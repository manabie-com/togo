projectList=(
  # "ms-account"
  # "ms-auth"
  # "ms-bill"
  # "ms-payment"
  # "ms-chat"
  # "ms-order"
  # "ms-customer"
  # "ms-document"
  # "ms-order"
  # "ms-exporting"
  # "ms-order"
  # "ms-invoice"
  # "ms-mail"
  # "ms-multi-platform"
  # "ms-notification"
  # "ms-order"
  # "ms-order"
  # "ms-permission"
  # "ms-product"
  # "ms-promotion"
  # "ms-property"
  # "ms-registration"
  # "ms-reporting"
  # "ms-order"
  # "ms-order"
  # "ms-order"
  # "ms-stock"
  # "ms-order"
  # "ms-order"
  # "ms-order"
  # "ms-system"
  # "ms-transactionHistory"
  "ms-productv2"
  # "ms-config"
)

module="user"

simple="task"
target="user"
simple2="Task"
target2="User"
firstUppercaseSimple="TransactionHistory"
firstUppercaseTarget="Task"
UppercaseSimple="USER_EMAIL"
UppercaseTarget="USER_EMAIL"



# for file in ${module}
#     do
#     #echo $file
#       sed -i -e "s/${simple}/${target}/g" $file
#       sed -i -e "s/${firstUppercaseSimple}/${firstUppercaseTarget}/g" $file
#       sed -i -e "s/${UppercaseSimple}/${UppercaseTarget}/g" $file
#       mv $file ${file//${simple}/${target}}
#       mv $file ${file//${firstUppercaseSimple}/${firstUppercaseTarget}}
#     done

for file in ${module}/*
    do
    #echo $file
      sed -i -e "s/${simple}/${target}/g" $file
      sed -i -e "s/${simple2}/${target2}/g" $file
      sed -i -e "s/${firstUppercaseSimple}/${firstUppercaseTarget}/g" $file
      sed -i -e "s/${UppercaseSimple}/${UppercaseTarget}/g" $file
      mv $file ${file//${simple}/${target}}
      mv $file ${file//${simple2}/${target2}}
      # mv $file ${file//${firstUppercaseSimple}/${firstUppercaseTarget}}
      # mv $file ${file//${UppercaseSimple}/${UppercaseSimple}}
    done

for file in ${module}/*/*
    do
    #echo $file
      sed -i -e "s/${simple}/${target}/g" $file
      sed -i -e "s/${simple2}/${target2}/g" $file
      sed -i -e "s/${firstUppercaseSimple}/${firstUppercaseTarget}/g" $file
      sed -i -e "s/${UppercaseSimple}/${UppercaseTarget}/g" $file
      mv $file ${file//${simple}/${target}}
      mv $file ${file//${simple2}/${target2}}
      # mv $file ${file//${firstUppercaseSimple}/${firstUppercaseTarget}}
      # mv $file ${file//${UppercaseSimple}/${UppercaseSimple}}
    done
for file in ${module}/*/*/*
    do
    #echo $file
      sed -i -e "s/${simple}/${target}/g" $file
      sed -i -e "s/${simple2}/${target2}/g" $file
      sed -i -e "s/${firstUppercaseSimple}/${firstUppercaseTarget}/g" $file
      sed -i -e "s/${UppercaseSimple}/${UppercaseTarget}/g" $file
      mv $file ${file//${simple}/${target}}
      mv $file ${file//${simple2}/${target2}}
      # mv $file ${file//${firstUppercaseSimple}/${firstUppercaseTarget}}
      # mv $file ${file//${UppercaseSimple}/${UppercaseSimple}}
    done
for file in ${module}/*/*/*/*
    do
    #echo $file
      sed -i -e "s/${simple}/${target}/g" $file
      sed -i -e "s/${simple2}/${target2}/g" $file
      sed -i -e "s/${firstUppercaseSimple}/${firstUppercaseTarget}/g" $file
      sed -i -e "s/${UppercaseSimple}/${UppercaseTarget}/g" $file
      mv $file ${file//${simple}/${target}}
      mv $file ${file//${simple2}/${target2}}
      # mv $file ${file//${firstUppercaseSimple}/${firstUppercaseTarget}}
      # mv $file ${file//${UppercaseSimple}/${UppercaseSimple}}
    done
for file in ${module}/*/*/*/*/*
    do
    #echo $file
      sed -i -e "s/${simple}/${target}/g" $file
      sed -i -e "s/${simple2}/${target2}/g" $file
      sed -i -e "s/${firstUppercaseSimple}/${firstUppercaseTarget}/g" $file
      sed -i -e "s/${UppercaseSimple}/${UppercaseTarget}/g" $file
      mv $file ${file//${simple}/${target}}
      mv $file ${file//${simple2}/${target2}}
      # mv $file ${file//${firstUppercaseSimple}/${firstUppercaseTarget}}
      # mv $file ${file//${UppercaseSimple}/${UppercaseSimple}}
    done