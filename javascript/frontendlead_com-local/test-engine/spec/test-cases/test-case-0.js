it('should verify the result is an array', (done) => {
    const result = flattenArray([1, [2, 3]]);
    expect(Array.isArray(result)).to.be.true;
    done();
  });
  