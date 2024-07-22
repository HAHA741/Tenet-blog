import { NodeSSH } from "node-ssh";
import archiver from "archiver";
import fs from "fs";
import sshConfig from "./sshConfig.mjs";

const config = {
  dirName: "TenetBlog",
  host: sshConfig.host,
  username: sshConfig.username,
  password: sshConfig.password,
  deployDir: "/home/Tenet/",
  localDir: process.cwd(),
};
//1.打包2.压缩打包文件3.连接服务器4.将压缩文件传入服务器5.服务器命令删除源文件6.服务器命令解压压缩文件
async function connectSSH() {
  try {
    const ssh = new NodeSSH();
    await ssh.connect({
      host: config.host,
      username: config.username,
      password: config.password,
    });
    console.log("连接成功");
    ssh.putFile(
      `${config.localDir}/${config.dirName}.zip`,
      `${config.deployDir}/${config.dirName}.zip`
    );
    console.log("文件上传成功");
    return ssh;
  } catch (err) {
    console.error("Error:", err);
  }
}
async function zipDist(sourceDir, outPath) {
  try {
    const output = fs.createWriteStream(outPath);
    const archive = archiver("zip", {
      zlib: { level: 9 },
    }).on("error", (err) => {
      throw err;
    });
    archive.on("error", (err) => {
      throw err;
    });

    // 通过管道方法将输出流存档到文件
    archive.pipe(output);

    // 从subdir子目录追加内容并重命名
    archive.directory(sourceDir, true);

    // 完成打包归档
    archive.finalize();
  } catch (err) {
    console.log(err, "err");
  }
}
async function runCommand(ssh, command, options) {
  let result = await ssh.execCommand(command, options);
  // if (result.stdout) {
  //   console.log('命令执行成功');
  // }
  return result;
}

//上传文件
async function uploadFile() {
  //压缩文件
  await zipDist(config.dirName, `${config.dirName}.zip`);
  //连接服务器
  let ssh = await connectSSH();
  let command1 = `cd ${config.deployDir}`;
  let command2 = `jar xvf ${config.dirName}.zip`;
  let command3 = `rm -f ${config.dirName}.zip`;
  await runCommand(ssh, command1, { cwd: config.deployDir });
  setTimeout(async () => {
    await runCommand(ssh, command2, { cwd: config.deployDir });
    await runCommand(ssh, command3, { cwd: config.deployDir });
    console.log("部署成功");
    process.exit(0);
  }, 1000);
}
uploadFile();
