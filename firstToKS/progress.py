def save_file(boy, girl, count):
    file_name_boy = 'boy' + str(count) + '.txt'
    file_name_girl = 'girl' + str(count) + '.txt'
    boy_file = open(file_name_boy, 'w')
    girl_file = open(file_name_girl, 'w')
    boy_file.writelines(boy)
    girl_file.writelines(girl)
    boy_file.close()
    girl_file.close()


def action_modify(file_name):
    f = open(file_name)
    line_num = 0
    arr=[]

    # for each_line in f:
    #     arr.append(each_line)
    #     line_num += 1
    #     print(each_line)

    count = 0

    for each_line in f:
        print(each_line.split('='))
        count += 1
    f.close()


action_modify("/root/analysisVerificationOfsystem/firstToKS/test")
str="x=y+1"

# x=7
# x=eval("x*3")
# print(x)
# x=eval("x*3")
# print(x)

x = 1
# var = "x"
varName = 'x'
locals()[varName] = 3
print(x)

