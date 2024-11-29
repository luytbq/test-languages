it('should handledeeply nested arrays', (done) => {
  const input = [[[[1]]], 2, [3, 4], 5];
  const expected = [1, 2, 3, 4];
  const result = flatternArray(input);
  expect(result).to.deep.equal(expected);
  done();
})