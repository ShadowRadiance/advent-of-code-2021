require "./cave/solver"

RSpec.describe Cave::Solver do
  let(:solver) { Cave::Solver.new(map) }
  let(:map) { Cave::Map.new(input)}
  let(:input) { Cave::Input.new(data) }

  context "with 3 node example" do
    let(:data) {
      <<~DATA
        start-A
        A-end
      DATA
    }
    it "is expected to find 1 paths" do
      expect(solver.count_paths).to eq(1)
    end
  end

  context "with 4 node example" do
    let(:data) {
      <<~DATA
        start-A
        start-b
        b-end
        A-end
      DATA
    }
    it "is expected to find 2 paths" do
      expect(solver.count_paths).to eq(2)
    end
  end

  context "with 6 node example" do
    let(:data) {
      <<~DATA
        start-A
        start-b
        A-c
        A-b
        b-d
        A-end
        b-end
      DATA
    }
    it "is expected to find 10 paths" do
      expect(solver.count_paths).to eq(10)
    end
  end

  context "with 8 node example" do
    let(:data) {
      <<~DATA
        dc-end
        HN-start
        start-kj
        dc-start
        dc-HN
        LN-dc
        HN-end
        kj-sa
        kj-HN
        kj-dc
      DATA
    }
    it "is expected to find 19 paths" do
      expect(solver.count_paths).to eq(19)
    end
  end

  context "with 11 node example" do
    let(:data) {
      <<~DATA
        fs-end
        he-DX
        fs-he
        start-DX
        pj-DX
        end-zg
        zg-sl
        zg-pj
        pj-he
        RW-he
        fs-DX
        pj-RW
        zg-RW
        start-pj
        he-WI
        zg-he
        pj-fs
        start-RW
      DATA
    }
    it "is expected to find 226 paths" do
      expect(solver.count_paths).to eq(226)
    end
  end

end