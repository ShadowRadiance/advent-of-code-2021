require "./cave/solver"

RSpec.describe Cave::Solver do
  let(:solver) { Cave::Solver.new(map, repeats: repeats) }
  let(:map) { Cave::Map.new(input)}
  let(:input) { Cave::Input.new(data) }
  let(:repeats) { 0 }


  context "with 3 node example" do
    let(:data) {
      <<~DATA
        start-A
        A-end
      DATA
    }
    context "no small repeats" do
      it "is expected to find 1 paths" do
        expect(solver.count_paths).to eq(1)
      end
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
    context "no small repeats" do
      it "is expected to find 2 paths" do
        expect(solver.count_paths).to eq(2)
      end
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
    context "no small repeats" do
      it "is expected to find 10 paths" do
        expect(solver.count_paths).to eq(10)
      end
    end

    context "one small repeat" do
      let(:repeats) { 1 }
      it "is expected to find 36 paths" do
        expect(solver.count_paths).to eq(36)
      end
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
    context "no small repeats" do
      it "is expected to find 19 paths" do
        expect(solver.count_paths).to eq(19)
      end
    end

    context "one small repeat" do
      let(:repeats) { 1 }
      it "is expected to find 103 paths" do
        expect(solver.count_paths).to eq(103)
      end
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
    context "no small repeats" do
      it "is expected to find 226 paths" do
        expect(solver.count_paths).to eq(226)
      end
    end

    context "one small repeat" do
      let(:repeats) { 1 }
      it "is expected to find 3509 paths" do
        expect(solver.count_paths).to eq(3509)
      end
    end
  end
end