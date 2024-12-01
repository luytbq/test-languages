it('should handle deeply nested arrays', (done) => {
    const input = [[[[[1]]]], 2, [[3], 4]];
    const expected = [1, 2, 3, 4];
    const result = flattenArray(input);
    expect(result).toEqual(expected);
    done();
  });
  