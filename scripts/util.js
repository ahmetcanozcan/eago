const handleErr = (err) => {
  if (err) {
    console.log(err);
    process.exit(1);
  }
};

module.exports = { handleErr };
