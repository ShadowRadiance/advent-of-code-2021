# frozen_string_literal: true

require "./image_enhancer"

RSpec.describe ImageEnhancer do
  subject { ImageEnhancer.new(algorithm) }
  let(:algorithm) { "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#" }
  let(:image) { InfiniteImage.new(data) }
  let(:data) {
    [
      "#..#.",
      "#....",
      "##..#",
      "..#..",
      "..###",
    ]
  }

  context "with a normal algorithm" do
    context "after 0 enhancements" do
      it "has the correct number of lights" do
        expect(image.count_lit).to eq(10)
      end
  
      it "can display the enhanced image" do
        expect(image.to_s).to eq(
          <<~IMAGE.chomp
          #..#.
          #....
          ##..#
          ..#..
          ..###
          IMAGE
        )
      end
    end
  
    context "after 1 enhancement" do
      before { @result = subject.enhance(image) }
  
      it "has the correct number of lights" do
        expect(@result.count_lit).to eq(24)
      end
  
      it "can display the enhanced image" do
        expect(@result.to_s).to eq(
          <<~IMAGE.chomp
          .##.##.
          #..#.#.
          ##.#..#
          ####..#
          .#..##.
          ..##..#
          ...#.#.
          IMAGE
        )
      end
    end
    
    context "after 2 enhancements" do
      before { @result = subject.enhance(subject.enhance(image)) }
      
      it "has the correct number of lights" do
        expect(@result.count_lit).to eq(35)
      end
      
      it "can display the enhanced image" do
        expect(@result.to_s).to eq(
          <<~IMAGE.chomp
          .......#.
          .#..#.#..
          #.#...###
          #...##.#.
          #.....#.#
          .#.#####.
          ..#.#####
          ...##.##.
          ....###..
          IMAGE
        )
      end
  
    end

    context "after 50 enhancements" do
      before {
        i = image
        50.times { i = subject.enhance(i) }
        @result = i
      }

      it "has the correct number of lights" do
        expect(@result.count_lit).to eq(3351)
      end
    end
  end

  context "with a background-flipping algorithm" do
    let(:data) { 
      [
        "..#",
        "#..",
        ".#.",
      ]
    }
    let(:algorithm) { "#.#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..." }

    context "after 0 enhancements" do
      it "has the correct number of lights" do
        expect(image.count_lit).to eq(3)
      end

      it "has the correct number of darks" do
        expect(image.count_unlit).to eq(Float::INFINITY)
      end
  
      it "can display the enhanced image" do
        expect(image.to_s).to eq(
          <<~IMAGE.chomp
          ..#
          #..
          .#.
          IMAGE
        )
      end
    end
  
    context "after 1 enhancement" do
      before { @result = subject.enhance(image) }
  
      it "has the correct number of lights" do
        expect(@result.count_lit).to eq(Float::INFINITY)
      end

      it "has the correct number of darks" do
        expect(@result.count_unlit).to eq(11)
      end
  
      it "can display the enhanced image" do
        expect(@result.to_s).to eq(
          <<~IMAGE.chomp
          ##.##
          .#..#
          ##...
          .#.##
          #.#.#
          IMAGE
        )
      end
    end
    
    context "after 2 enhancements" do
      before { @result = subject.enhance(subject.enhance(image)) }
      
      it "has the correct number of lights" do
        expect(@result.count_lit).to eq(15)
      end
      
      it "has the correct number of darks" do
        expect(@result.count_unlit).to eq(Float::INFINITY)
      end

      it "can display the enhanced image" do
        expect(@result.to_s).to eq(
          <<~IMAGE.chomp
          ..#....
          .#.##..
          ##...##
          .....##
          ##...#.
          .#...#.
          IMAGE
        )
      end
  
    end

  end
end
